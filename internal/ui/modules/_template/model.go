package moduleTemplate

import (

)

type Model struct {
	data int
}

func NewModel() *Model {
	m := &Model{
		data: 42
	}

	return m
}

// ----- Data setters -----
func (m *Model) SetData(data int) {
	m.data = data
}

// ----- Data getters -----
func (m *Model) GetData() int {
	return m.data
}
