package graph

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
)

// Edge represents an edge in a graph
type Edge interface {
	From() Vertex
	To() Vertex

	Flow() float32
	SetFlow(float32)

	Capacity() float32

	Residual() float32
}

// Vertex represents a vertex in a graph
type Vertex interface {
	ID() int

	Height() int
	SetHeight(int)

	Excess() float32
	SetExcess(float32)

	Edges() []Edge
	Edge(int) Edge
	NumEdges() int

	IsNeighbor(int) bool
}

// Graph represents a network graph
type Graph interface {
	Source() Vertex
	SetSource(int)

	Sink() Vertex
	SetSink(int)

	Vertices() []Vertex
	NumVertices() int
	AddVertex() int
	GetVertex(int) Vertex
	RemoveVertex(int)

	Edges() []Edge
	NumEdges() int
	AddEdge(int, int, float32, float32) bool
	GetEdge(int, int) Edge
	RemoveEdge(int, int)
}

// New constructs a new graph
func New() Graph {
	return &graph{
		source:  -1,
		sink:    -1,
		mapping: make(map[int]wrapper),
	}
}

// Properties are the properties a graph might have
type Properties struct {
	SelfLoops    bool
	FlowGraph    bool
	ReverseEdges bool
}

// Random generates a random graph given the number
// of vertices, edges, and a seed.  If the parameters
// are invalid, i.e negative values, the value nil is
// returned
func Random(vertices, edges int, seed int64, prop Properties) Graph {
	if vertices < 0 || edges < 0 || edges < vertices-1 {
		return nil
	}

	rand.Seed(seed)

	result := New()
	for i := 0; i < vertices; i++ {
		result.AddVertex()
	}

	src := -1
	snk := -1

	if prop.FlowGraph {

		for src == snk {
			src = rand.Intn(vertices)
			snk = rand.Intn(vertices)
		}

		result.SetSource(src)
		result.SetSink(snk)

		prop.ReverseEdges = false
		prop.SelfLoops = false
	}

	for i := 0; i < edges; i++ {

		for {

			from := rand.Intn(vertices)
			to := rand.Intn(vertices)

			add := true

			if !prop.SelfLoops && from == to {
				add = false
			} else if !prop.ReverseEdges && result.GetEdge(to, from) != nil {
				add = false
			} else if prop.FlowGraph && (to == src || from == snk) {
				add = false
			}

			if add && result.GetEdge(from, to) == nil {
				result.AddEdge(from, to, 0, rand.Float32())
				break
			}
		}
	}

	return result
}

// WriteGraph takes a file and graph which then writes
// the graph to the file.  This format is similar to
// the DIMACS graph format but differs on the amount of
// information given for each vertex
func WriteGraph(filename string, g Graph) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	var buffer bytes.Buffer

	buffer.WriteString("c\n")
	buffer.WriteString("c File written using jeff/ds/graph/WriteGraph\n")
	buffer.WriteString("c\n")
	buffer.WriteString("p \n")

	src := g.Source()
	snk := g.Sink()

	if g.Source() != nil {
		buffer.WriteString(fmt.Sprintf("p \t max \t %d \t %d\n", g.NumVertices(), g.NumEdges()))
		buffer.WriteString(fmt.Sprintf("n \t %d \t s\n", src.ID()))
		buffer.WriteString(fmt.Sprintf("n \t %d \t t\n", snk.ID()))
	} else {
		buffer.WriteString(fmt.Sprintf("p none %d %d\n", g.NumVertices(), g.NumEdges()))
	}

	for _, v := range g.Vertices() {
		for _, e := range v.Edges() {
			buffer.WriteString(fmt.Sprintf("a \t %d \t %d \t %f \t %f\n", e.From().ID(), e.To().ID(), e.Flow(), e.Capacity()))
		}
	}

	file.Truncate(0)
	n, err := file.Write(buffer.Bytes())
	if err != nil {
		return err
	}
	fmt.Println("Bytes written:", n)
	file.Sync()

	return nil
}

// LoadGraph will read a file which has a graph written
// to by WriteGraph
func LoadGraph(filename string) (Graph, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	input := bufio.NewReader(file)

	result := New().(*graph)

	for {
		if line, err := input.ReadString('\n'); err == nil {
			switch line[0] {
			case 'a':
				from := 0
				to := 0
				flow := float32(0)
				cap := float32(0)
				
				if n, err := fmt.Sscanf(line, "a \t %d \t %d \t %f \t %f\n", &from, &to, &flow, &cap); err != nil {
					return nil, err
				} else if n != 4 {
					return nil, fmt.Errorf("Expected 4 items, read only %d", n)
				}

				if result.GetVertex(from) == nil {
					result.addVertex(from)
				}

				if result.GetVertex(to) == nil {
					result.addVertex(to)
				}

				result.AddEdge(from, to, flow, cap)
			case 'c':
			case 'n':
				id := 0
				which := byte(0)
				fmt.Sscanf(line,"n \t %d \t %c\n", &id, &which)
				switch which {
				case 's':
					result.SetSource(id)
				case 't':
					result.SetSink(id)
				}
			default:
				// Ignore?
			}
		} else if err != io.EOF {
			return nil, err
		} else {
			break
		}
	}

	return result, nil
}
