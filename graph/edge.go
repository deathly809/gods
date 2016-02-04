package graph

type edge struct {
	from, to       *vertex
	flow, capacity float32
}

func (e *edge) From() Vertex {
	return e.from
}

func (e *edge) To() Vertex {
	return e.to
}

func (e *edge) Flow() float32 {
	return e.flow
}

func (e *edge) Capacity() float32 {
	return e.capacity
}

func (e *edge) SetFlow(flow float32) {
	e.flow = flow
}

func (e *edge) Residual() float32 {
	return e.Capacity() - e.Flow()
}
