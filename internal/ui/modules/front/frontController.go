package moduleTemplate

import (
	"fyne.io/fyne/v2"
	"hrubos.dev/collectorsden/internal/database"
)

type Controller struct {
	*Model
	*View
	window fyne.Window
	app fyne.App
}

func NewController(m *Model, v *View, app fyne.App, window fyne.Window) *Controller {
	c := &Controller{
		Model: m,
		View: v,
		window: window,
		app: app,
	}

	c.View.BindList(
		c.getSingleCategoryFromSlice,
		c.getLengthOfSlice,
	)

	return c
}

//TODO completely rewrite this mess later
func (c *Controller) getSingleCategoryFromSlice(index int) database.Category {
	categories := c.Model.GetCategories()
	return categories[index]
}

func (c *Controller) getLengthOfSlice() int {
	categories := c.Model.GetCategories()
	return len(categories)
}
