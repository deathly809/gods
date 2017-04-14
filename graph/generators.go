package graph

import (
	"fmt"
	"math/rand"
)

// Random generates a random graph given the number
// of vertices, edges, and a seed.  If the parameters
// are invalid, i.e negative values, the value nil is
// returned
func Random(vertices, edges int, seed int64, prop Properties) (Graph, error) {
	if vertices < 0 || edges < 0 || edges < vertices-1 {
		return nil, fmt.Errorf("invalid parameters: #vertices %v, #edges %v", vertices, edges)
	}
	rand.Seed(seed)
	result := New(prop)

	for i := 0; i < vertices; i++ {
		result.AddVertex()
	}

	for i := 0; i < edges; i++ {

		for {

			from, _ := result.GetVertex(rand.Intn(vertices))
			to, _ := result.GetVertex(rand.Intn(vertices))

			edge, err := result.AddEdge(from, to)
			if err != nil || edge == nil {
				continue
			}
			break
		}
	}
	return result, nil
}
