package moduleTemplate

import (

)

type Model struct {
	data int
}

func NewModel() *Model {
	fm := &Model{
		data = 42
	}

	return fm
}

// ----- Data setters -----
func (fm *Model) SetData(data int) {
	fm.data = data
}

// ----- Data getters -----
func (fm *Model) GetData() int {
	return fm.data
}
