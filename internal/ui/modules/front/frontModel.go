package moduleTemplate

import (
	"hrubos.dev/collectorsden/internal/database"
)

type Model struct {
	categories []database.Category
}

func NewModel() *Model {
	m := &Model{
		categories: make([]database.Category, 0),
	}

	m.categories = database.AllCategories()

	return m
}

// ----- Data setters -----
func (m *Model) SetCategories(categories []database.Category) {
	m.categories = categories
}

// ----- Data getters -----
func (m *Model) GetCategories() []database.Category {
	return m.categories
}
