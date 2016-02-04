package graph

import (
	"math/rand"
	"testing"
)

const (
	N        = 1000
	M        = N * 10
	ToRemove = N / 10
	After    = M - ToRemove
	File     = "derp.max"
)

func TestNewGraph(t *testing.T) {
	g := New()
	if g == nil {
		t.Fatal("The graph is nil")
	}

	if g.NumEdges() != 0 {
		t.Fatal("The empty graph reports edges exist")
	}

	if g.NumVertices() != 0 {
		t.Fatal("The empty graph reports vertices exist")
	}
}

func TestGraphInsertVertices(t *testing.T) {
	g := New()

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	if g.NumEdges() != 0 {
		t.Fatal("The graph with no edges reports edges exist")
	}

	if g.NumVertices() != N {
		t.Fatalf("Expected %d vertoces, but found %d vertices", N, g.NumVertices())
	}
}
func TestGraphInsertEdges(t *testing.T) {
	g := New()

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	for i := 0; i < M; i++ {
		for {
			from := rand.Intn(N)
			to := rand.Intn(N)
			if g.AddEdge(from, (to)%N, 0, 0) {
				break
			}
		}
	}

	if g.NumEdges() != M {
		t.Fatalf("Expected %d edges, but found %d edges", M, g.NumEdges())
	}

	if g.NumVertices() != N {
		t.Fatalf("Expected %d vertoces, but found %d vertices", N, g.NumVertices())
	}
}

type intPair struct {
	a, b int
}

func TestGraphRemoveEdge(t *testing.T) {
	g := New()

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	var edges []intPair

	for i := 0; i < M; i++ {

		for {
			from := rand.Intn(N)
			to := rand.Intn(N)
			if g.AddEdge(from, (to)%N, 0, 0) {
				edges = append(edges, intPair{from, to})
				break
			}
		}
	}

	if len(edges) != M {
		t.Fatalf("Number of edges should be %d but %d was reported", len(edges), M)
	}

	for i := 0; i < ToRemove; i++ {
		e := edges[len(edges)-1]
		edges = edges[:len(edges)-1]
		g.RemoveEdge(e.a, e.b)
	}

	if g.NumEdges() != After {
		t.Fatalf("Expected %d edges, but found %d edges", After, g.NumEdges())
	}

	for _, tuple := range edges {
		if g.GetEdge(tuple.a, tuple.b) == nil {
			t.Fatalf("Edge (%d,%d) should be connected, but is not", tuple.a, tuple.b)
		}
	}

	if g.NumVertices() != N {
		t.Fatalf("Expected %d vertoces, but found %d vertices", N, g.NumVertices())
	}
}

func TestBFS(t *testing.T) {
	N := 100000
	g := New()

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	g.AddEdge(0, 1, 0, 0)
	g.AddEdge(0, 2, 0, 0)
	g.AddEdge(0, 3, 0, 0)
	g.AddEdge(0, 4, 0, 0)
	g.AddEdge(0, 5, 0, 0)
	g.AddEdge(0, 6, 0, 0)
	g.AddEdge(0, 7, 0, 0)

	res := BFS(g)

	if len(res) != N {
		t.Errorf("Expected %d results but found %d results", g.NumVertices(), len(res))
	}
}

func TestRandom(t *testing.T) {

	r := Random(N, M, 1022, Properties{FlowGraph: true})

	if r.NumVertices() != N {
		t.Errorf("Expected %d vertices but found %d", N, r.NumVertices())
	}

	if r.NumEdges() != M {
		t.Errorf("Expected %d edges but found %d", M, r.NumEdges())
	}

	src := r.Source()
	snk := r.Sink()

	if snk.NumEdges() != 0 {
		t.Fatal("Sink has outgoing edges!")
	}

	for i := 0; i < N; i++ {
		if r.GetEdge(i, src.ID()) != nil {
			t.Fatal("Source has incoming edge")
		}
	}
}

func TestWriter(t *testing.T) {

	r := Random(N, M, 1, Properties{FlowGraph: true})
	if err := WriteGraph(File, r); err != nil {
		t.Fatal(err.Error())
	}
	g, err := LoadGraph(File)
	if err != nil {
		t.Fatal(err.Error())
	}

	if g.NumVertices() != N {
		t.Errorf("Expected %d vertices but found %d", N, g.NumVertices())
	}

	if g.NumEdges() != M {
		t.Errorf("Expected %d vertices but found %d", M, g.NumEdges())
	}
}
