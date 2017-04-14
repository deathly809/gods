package graph

import "encoding/json"

// Marshable means we can convert to and from a JSON object
type Marshable interface {
	json.Marshaler
	json.Unmarshaler
}

// Edge represents an edge in a graph
type Edge interface {
	Marshable

	From() Vertex
	To() Vertex

	Get() Marshable
	Set(Marshable)
}

// Vertex represents a vertex in a graph
type Vertex interface {
	Marshable

	ID() int
	Get() Marshable
	Set(Marshable)
}

// Graph represents a network graph
type Graph interface {
	Marshable

	// Vertex Methods
	AddVertex() Vertex
	GetVertex(int) (Vertex, error)
	RemoveVertex(Vertex) error
	Vertices() <-chan Vertex
	Neighbors(Vertex) <-chan Edge
	NumVertices() int

	// Edge Methods
	AddEdge(Vertex, Vertex) (Edge, error)
	GetEdge(Vertex, Vertex) (Edge, error)
	RemoveEdge(Edge) error
	Edges() <-chan Edge
	NumEdges() int
}

// Properties are the properties a graph might have
type Properties struct {
	SelfLoops bool
	Directed  bool
}
