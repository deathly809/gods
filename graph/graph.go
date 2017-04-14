package graph

import (
	"encoding/json"
	"fmt"
)

const (
	vertexDoesNotBelongMsg = "vertex does not belong to this graph"
	edgeDoesNotBelongMsg   = "edge does not belong to this graph"
	edgeDoesNotExistMsg    = "edge does not exist"
	vertexDoesNotExistMsg  = "vertex does not exist"
	selfLoopNotAllowedMsg  = "self-loops are not allowed"
	getEdgeErrorMsg        = "could not return edge"
	couldNotAddEdgeMsg     = "could not add edge to graph"
	couldNotRemoveEdgeMsg  = "could not remove edge from graph"
	couldNotGetEdgeMsg     = "could not get edge in graph"
)

type graph struct {
	VertexID  int
	GraphType Properties
	vertices  map[int]*vertex
	edges     map[int]map[int]*edge
	numEdges  int
}

func (g *graph) Directed() bool {
	return g.GraphType.Directed
}

func (g *graph) isMyVertex(u Vertex) bool {
	if v, err := g.GetVertex(u.ID()); err == nil {
		return v == u
	}
	return false
}

// Validate the vertices and make sure they are in their correct order
func (g *graph) prepareVertices(from, to Vertex) (Vertex, Vertex, error) {
	fromIsBad := !g.isMyVertex(from)
	toIsBad := !g.isMyVertex(to)

	if fromIsBad || toIsBad {
		return from, to, fmt.Errorf("%s : %v : %v, %v : %v",
			vertexDoesNotBelongMsg,
			from.ID(),
			fromIsBad,
			to.ID(),
			toIsBad,
		)
	}

	if !g.GraphType.SelfLoops && from.ID() == to.ID() {
		return from, to, fmt.Errorf("%s", selfLoopNotAllowedMsg)
	}

	if g.Directed() && to.ID() < from.ID() {
		return from, to, nil
	}
	return to, from, nil
}

// 	Add an edge to the graph if it does not already exist
// 	If the new edge conflicts with the graph properties
// 	we return an error.
//	If the edge already exists we return nil
func (g *graph) AddEdge(from, to Vertex) (Edge, error) {

	to, from, err := g.prepareVertices(to, from)
	if err != nil {
		return nil, fmt.Errorf("%s : %v", couldNotAddEdgeMsg, err)
	}

	eList, _ := g.edges[from.ID()]
	if _, ok := eList[to.ID()]; ok {
		return nil, nil
	}

	nEdge := newEdge(g, from, to, nil)
	g.numEdges++
	eList[to.ID()] = nEdge
	return nEdge, nil
}

func (g *graph) RemoveEdge(e Edge) error {
	to, from, err := g.prepareVertices(e.From(), e.To())
	if err != nil {
		return fmt.Errorf("%s : %v", couldNotRemoveEdgeMsg, err)
	}

	if _, err := g.GetEdge(from, to); err != nil {
		return err
	}

	g.numEdges--
	delete(g.edges[from.ID()], to.ID())
	return nil
}

func (g *graph) GetEdge(from, to Vertex) (Edge, error) {
	to, from, err := g.prepareVertices(from, to)
	if err != nil {
		return nil, fmt.Errorf("%s : %v", couldNotGetEdgeMsg, err)
	}

	if edges, ok := g.edges[from.ID()]; ok {
		if e, ok := edges[to.ID()]; ok {
			return e, nil
		}
	}
	return nil, fmt.Errorf("%s", couldNotGetEdgeMsg)
}

func (g *graph) NumEdges() int {
	return g.numEdges
}

func (g *graph) Edges() <-chan Edge {
	result := make(chan Edge)
	go func() {
		defer close(result)
		for _, edgesForVertex := range g.edges {
			for _, edge := range edgesForVertex {
				result <- edge
			}
		}

	}()
	return result
}

// Vertex Functions
func (g *graph) AddVertex() Vertex {
	g.VertexID++

	g.edges[g.VertexID] = make(map[int]*edge)
	g.vertices[g.VertexID] = newVertex(g.VertexID, nil)

	return g.vertices[g.VertexID]
}
func (g *graph) RemoveVertex(v Vertex) error {
	if !g.isMyVertex(v) {
		return fmt.Errorf("%s : %v", vertexDoesNotBelongMsg, v.ID())
	}

	for k := range g.vertices {
		delete(g.edges[k], v.ID())
	}

	delete(g.vertices, v.ID())
	delete(g.edges, v.ID())

	for k, e := range g.edges {
		if !g.Directed() && k > v.ID() {
			break
		}
		delete(e, v.ID())
	}

	return nil
}

func (g *graph) GetVertex(id int) (Vertex, error) {
	if v, ok := g.vertices[id]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("%s", vertexDoesNotExistMsg)
}

func (g *graph) NumVertices() int {
	return len(g.vertices)
}

func (g *graph) Vertices() <-chan Vertex {
	result := make(chan Vertex)
	go func() {
		defer close(result)
		for _, vertex := range g.vertices {
			result <- vertex
		}
	}()
	return result
}

func (g *graph) Neighbors(v Vertex) <-chan Edge {
	if !g.isMyVertex(v) {
		return nil
	}
	result := make(chan Edge)
	go func() {
		defer close(result)
		if g.Directed() {
			for k, vertexEdges := range g.edges {
				if k == v.ID() {
					for _, e := range vertexEdges {
						result <- e
					}
				} else {
					if e, ok := vertexEdges[v.ID()]; ok {
						result <- e
					}
				}
			}
		} else {
			for k, vertexEdges := range g.edges {
				if k == v.ID() {
					for _, e := range vertexEdges {
						result <- e
					}
					return
				}
				if e, ok := vertexEdges[v.ID()]; ok {
					result <- e
				}
			}
		}

	}()
	return result
}

type wrapper struct {
	VertexID  int
	GraphType Properties
	Vertices  map[int]*vertex
	Edges     map[int]map[int]*edge
	NumEdges  int
}

func (g *graph) MarshalJSON() ([]byte, error) {
	wrap := wrapper{}
	wrap.VertexID = g.VertexID
	wrap.GraphType = g.GraphType
	wrap.Vertices = g.vertices
	wrap.Edges = g.edges
	wrap.NumEdges = g.numEdges
	return json.Marshal(wrap)
}

func (g *graph) UnmarshalJSON(d []byte) error {
	wrap := wrapper{}
	err := json.Unmarshal(d, &wrap)

	if err == nil {
		g.VertexID = wrap.VertexID
		g.GraphType = wrap.GraphType
		g.vertices = wrap.Vertices
		g.edges = wrap.Edges
		g.numEdges = wrap.NumEdges
	}
	return err
}

// New graph is returned
func New(t Properties) Graph {
	return &graph{
		GraphType: t,
		VertexID:  -1,
		edges:     make(map[int]map[int]*edge),
		vertices:  make(map[int]*vertex),
	}
}
