package graph

import (
	"encoding/json"
	"math/rand"
	"testing"
)

const (
	N        = 20
	M        = N * 5
	ToRemove = N / 10
	After    = M - ToRemove
	File     = "temporary.max"
)

func TestNewGraph(t *testing.T) {
	g := New(Properties{})
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
	g := New(Properties{})

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	if g.NumEdges() != 0 {
		t.Fatal("The graph with no edges reports edges exist")
	}

	if g.NumVertices() != N {
		t.Fatalf("Expected %d vertices, but found %d vertices", N, g.NumVertices())
	}
}
func TestGraphInsertEdges(t *testing.T) {
	g := New(Properties{Directed: true})

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	for i := 0; i < M; i++ {
		for {

			from, err := g.GetVertex(rand.Intn(N))
			if err != nil {
				t.Fatal(err)
			}

			to, err := g.GetVertex(rand.Intn(N))
			if err != nil {
				t.Fatal(err)
			}

			if to == from {
				continue
			}

			if _, err = g.GetEdge(from, to); err == nil {
				continue
			}

			if _, err = g.AddEdge(from, to); err != nil {
				continue
			}
			break
		}
	}

	if g.NumEdges() != M {
		t.Fatalf("Expected %d edges, but found %d edges", M, g.NumEdges())
	}

	if g.NumVertices() != N {
		t.Fatalf("Expected %d vertices, but found %d vertices", N, g.NumVertices())
	}
}

type intPair struct {
	a, b int
}

func _TestGraphRemoveEdge(t *testing.T) {
	g := New(Properties{})

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	var edges []Edge

	for i := 0; i < M; i++ {

		for {
			from, _ := g.GetVertex(rand.Intn(N))
			to, _ := g.GetVertex(rand.Intn(N))
			if edge, err := g.AddEdge(from, to); err == nil && edge != nil {
				edges = append(edges, edge)
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
		g.RemoveEdge(e)
	}

	if g.NumEdges() != After {
		t.Fatalf("Expected %d edges, but found %d edges", After, g.NumEdges())
	}

	for _, e := range edges {
		if _, err := g.GetEdge(e.From(), e.To()); err != nil {
			t.Fatalf("Edge (%d,%d) should be connected, but is not", e.From().ID(), e.To().ID())
		}
	}

	if g.NumVertices() != N {
		t.Fatalf("Expected %d vertices, but found %d vertices", N, g.NumVertices())
	}
}

func _TestBFS(t *testing.T) {
	N := 100000
	g := New(Properties{})

	for i := 0; i < N; i++ {
		g.AddVertex()
	}

	for i := 1; i < 8; i++ {
		from, err := g.GetVertex(0)
		if err != nil {
			t.Error(err)
		}
		to, err := g.GetVertex(i)
		if err != nil {
			t.Error(err)
		}

		_, err = g.AddEdge(from, to)

		if err != nil {
			t.Error(err)
		}
	}

	res := BFS(g)

	if len(res) != N {
		t.Errorf("Expected %d results but found %d results", g.NumVertices(), len(res))
	}
}

func TestRandom(t *testing.T) {

	r, err := Random(N, M, 1022, Properties{})
	if err != nil {
		t.Error(err)
	}

	if r.NumVertices() != N {
		t.Errorf("Expected %d vertices but found %d", N, r.NumVertices())
	}

	if r.NumEdges() != M {
		t.Errorf("Expected %d edges but found %d", M, r.NumEdges())
	}

}

func TestWriter(t *testing.T) {

	r, err := Random(N, M, 1, Properties{})
	if err != nil {
		t.Error(err)
	}

	data, err := json.Marshal(r)

	if err != nil {
		t.Error(err)
	}

	g := New(Properties{})
	err = json.Unmarshal(data, g)
	if err != nil {
		t.Error(err)
	}

	if g.NumVertices() != N {
		t.Errorf("Expected %d vertices but found %d", N, g.NumVertices())
	}

	if g.NumEdges() != M {
		t.Errorf("Expected %d edges but found %d", M, g.NumEdges())
	}
}
