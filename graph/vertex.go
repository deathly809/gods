package graph

type edgeSlice []Edge

func (e edgeSlice) Swap(i, j int) {
	tmp := e[i]
	e[i] = e[j]
	e[j] = tmp
}

type vertex struct {
	g          *graph
	id, height int
	excess     float32
	edges      edgeSlice
	mappings   map[int]wrapper
}

func (v *vertex) ID() int {
	return v.id
}

func (v *vertex) Height() int {
	return v.height
}

func (v *vertex) SetHeight(height int) {
	v.height = height
}

func (v *vertex) Excess() float32 {
	return v.excess
}

func (v *vertex) SetExcess(excess float32) {
	v.excess = excess
}

func (v *vertex) NumEdges() int {
	return len(v.edges)
}

func (v *vertex) Edges() []Edge {
	return v.edges
}

func (v *vertex) Edge(to int) Edge {
	wrp, present := v.mappings[to]
	if present {
		return wrp.value.(*edge)
	}
	return nil
}

func (v *vertex) connect(to int, flow, capacity float32) bool {
	e := v.Edge(to)
	if e != nil {
		return false
	}

	t := v.g.GetVertex(to)
	if t == nil {
		return false
	}

	e = &edge{
		from:     v,
		to:       t.(*vertex),
		flow:     flow,
		capacity: capacity,
	}

	v.g.edges = nil
	v.edges = append(v.edges, e)

	v.mappings[to] = wrapper{e, len(v.edges) - 1}

	return true
}

func (v *vertex) removeEdge(to int) {
	wrapper, present := v.mappings[to]
	if present {
		v.edges.Swap(wrapper.idx, len(v.edges)-1)
		v.edges = v.edges[:len(v.edges)-1]
		delete(v.mappings, to)
	}
}

func (v *vertex) IsNeighbor(to int) bool {
	_, present := v.mappings[to]
	return present
}
