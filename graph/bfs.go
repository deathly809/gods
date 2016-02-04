package graph

import "jeff/ds/queue"

type vResult struct {
	v     Vertex
	start int
	end   int
}

// SearchResult wraps the results for for Graph searches
type SearchResult map[Vertex]*vResult

func bfs(v Vertex, g Graph, res SearchResult, tick int) int {
	Q := queue.New()

	res[v] = &vResult{start: tick, v: v}
	Q.Enqueue(res[v])
	tick++

	for Q.Count() > 0 {
		s := Q.Dequeue().(*vResult)
		s.end = tick
		tick++
		for _, e := range s.v.Edges() {
			v = e.To()
			_, e := res[v]
			if !e {
				res[v] = &vResult{start: tick, v: v}
				Q.Enqueue(res[v])
				tick++
			}
		}
	}
	return tick
}

// BFS performs a breadth first search on a Graph
func BFS(g Graph) SearchResult {

	tick := 0

	result := SearchResult(make(map[Vertex]*vResult))

	for _, v := range g.Vertices() {
		if _, exists := result[v]; !exists {
			tick = bfs(v, g, result, tick)
		}
	}

	return result
}
