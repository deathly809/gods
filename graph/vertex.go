package graph

import (
	"encoding/json"
	"fmt"
)

type vertex struct {
	id   int
	data Marshable
}

func (v *vertex) ID() int {
	return v.id
}

func (v *vertex) Get() Marshable {
	return v.data
}

func (v *vertex) Set(data Marshable) {
	v.data = data
}

type vWrapper struct {
	ID   int
	Data Marshable
}

func (v *vertex) MarshalJSON() ([]byte, error) {
	wrapper := vWrapper{
		v.id,
		v.data,
	}
	return json.Marshal(wrapper)
}

func (v *vertex) UnmarshalJSON(data []byte) error {
	wrapper := vWrapper{}

	err := json.Unmarshal(data, &wrapper)
	if err == nil {
		v.id = wrapper.ID
		v.data = wrapper.Data
	} else {
		fmt.Println(err)
	}
	return err
}

func newVertex(id int, data Marshable) *vertex {
	return &vertex{
		id:   id,
		data: data,
	}
}
