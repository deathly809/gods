package graph

import "jeff/math"

type wrapper struct {
	value interface{}
	idx   int
}

type graph struct {
	id           int
	numEdges     int
	source, sink int
	vertices     []Vertex
	mapping      map[int]wrapper
	edges        []Edge
}

// Constructs new slice each time, use sparingly
func (g *graph) Edges() []Edge {
	if g.edges == nil {
		g.edges = make([]Edge, g.NumEdges())
		for _, v := range g.vertices {
			g.edges = append(g.edges, v.Edges()...)
		}
	}
	return g.edges
}

func (g *graph) Vertices() []Vertex {
	return g.vertices
}

func (g *graph) NumVertices() int {
	return len(g.vertices)
}

func (g *graph) NumEdges() int {
	return g.numEdges
}

func (g *graph) Source() Vertex {
	return g.GetVertex(g.source)
}

func (g *graph) Sink() Vertex {
	return g.GetVertex(g.sink)
}

func (g *graph) SetSource(source int) {
	g.source = source
}

func (g *graph) SetSink(sink int) {
	g.sink = sink
}

func (g *graph) GetEdge(from, to int) Edge {
	v := g.GetVertex(from)
	if v != nil {
		return v.Edge(to)
	}
	return nil
}

func (g *graph) RemoveEdge(from, to int) {
	v := g.GetVertex(from)
	if v != nil {
		g.numEdges--
		g.edges = nil
		v.(*vertex).removeEdge(to)
	}
}

func (g *graph) RemoveVertex(vid int) {
	v, exists := g.mapping[vid]
	if exists {

		N := g.NumVertices()
		r := g.vertices[N-1]

		g.vertices[v.idx] = r
		v, _ = g.mapping[r.ID()]
		v.idx = vid

		delete(g.mapping, vid)

		g.vertices = g.vertices[:N-1]

	}
}

func (g *graph) GetVertex(vid int) Vertex {
	v, exists := g.mapping[vid]
	if exists {
		return v.value.(*vertex)
	}
	return nil
}

func (g *graph) AddEdge(from, to int, flow, capacity float32) bool {

	v := g.GetVertex(from)
	if v != nil {

		if v.(*vertex).connect(to, flow, capacity) {
			g.edges = nil
			g.numEdges++
			return true
		}
	}
	return false
}

// used for building graphs from scratch
func (g *graph) addVertex(id int) int {
	v := &vertex{
		id:       id,
		excess:   0,
		height:   0,
		g:        g,
		mappings: make(map[int]wrapper),
	}
	g.mapping[v.id] = wrapper{value: v, idx: g.NumVertices()}
	g.vertices = append(g.vertices, v)
	g.id = math.MaxInt(g.id, id+1)
	return id
}

func (g *graph) AddVertex() int {
	return g.addVertex(g.id)
}
