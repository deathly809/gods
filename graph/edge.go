package graph

import (
	"encoding/json"
)

type edge struct {
	g    Graph
	from int
	to   int
	data Marshable
}

func (e *edge) From() Vertex {
	v, _ := e.g.GetVertex(e.from)
	return v
}

func (e *edge) To() Vertex {
	v, _ := e.g.GetVertex(e.to)
	return v
}

func (e *edge) Get() Marshable {
	return e.data
}

func (e *edge) Set(data Marshable) {
	e.data = data
}

type eWrapper struct {
	From int
	To   int
	Data Marshable
}

func (e *edge) MarshalJSON() ([]byte, error) {
	wrapper := eWrapper{
		e.from,
		e.to,
		e.data,
	}
	return json.Marshal(wrapper)
}

func (e *edge) UnmarshalJSON(data []byte) error {
	wrapper := eWrapper{}
	err := json.Unmarshal(data, &wrapper)
	if err == nil {
		e.from = wrapper.From
		e.to = wrapper.To
		e.data = wrapper.Data
	}
	return err
}

func newEdge(g Graph, from, to Vertex, data Marshable) *edge {
	return &edge{
		g:    g,
		from: from.ID(),
		to:   to.ID(),
		data: data,
	}
}
