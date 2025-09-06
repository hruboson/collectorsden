package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// ----- 1/3 - 2/3 -----
type OneThirdTwoThirdLayout struct{}

func (l OneThirdTwoThirdLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 2 {
		return
	}
	oneThird := size.Width / 3
	objects[0].Resize(fyne.NewSize(oneThird, size.Height))
	objects[0].Move(fyne.NewPos(0, 0))

	twoThird := size.Width - oneThird
	objects[1].Resize(fyne.NewSize(twoThird, size.Height))
	objects[1].Move(fyne.NewPos(oneThird, 0))
}

func (l OneThirdTwoThirdLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minW := fyne.Max(objects[0].MinSize().Width*3, objects[1].MinSize().Width*3/2)
	minH := fyne.Max(objects[0].MinSize().Height, objects[1].MinSize().Height)
	return fyne.NewSize(minW, minH)
}

// ----- 2/3 - 1/3 -----
type TwoThirdOneThirdLayout struct{}

func (l TwoThirdOneThirdLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 2 {
		return
	}
	twoThird := 2 * size.Width / 3
	objects[0].Resize(fyne.NewSize(twoThird, size.Height))
	objects[0].Move(fyne.NewPos(0, 0))

	oneThird := size.Width - twoThird
	objects[1].Resize(fyne.NewSize(oneThird, size.Height))
	objects[1].Move(fyne.NewPos(twoThird, 0))
}

func (l TwoThirdOneThirdLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minW := fyne.Max(objects[0].MinSize().Width*3/2, objects[1].MinSize().Width*3)
	minH := fyne.Max(objects[0].MinSize().Height, objects[1].MinSize().Height)
	return fyne.NewSize(minW, minH)
}

// ----- 1/4 - 3/4 -----
type OneFourthThreeFourthLayout struct{}

func (l OneFourthThreeFourthLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 2 {
		return
	}
	oneFourth := size.Width / 4
	objects[0].Resize(fyne.NewSize(oneFourth, size.Height))
	objects[0].Move(fyne.NewPos(0, 0))

	threeFourth := size.Width - oneFourth
	objects[1].Resize(fyne.NewSize(threeFourth, size.Height))
	objects[1].Move(fyne.NewPos(oneFourth, 0))
}

func (l OneFourthThreeFourthLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minW := fyne.Max(objects[0].MinSize().Width*4, objects[1].MinSize().Width*4/3)
	minH := fyne.Max(objects[0].MinSize().Height, objects[1].MinSize().Height)
	return fyne.NewSize(minW, minH)
}

// ----- 3/4 - 1/4 -----
type ThreeFourthOneFourthLayout struct{}

func (l ThreeFourthOneFourthLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	if len(objects) < 2 {
		return
	}
	threeFourth := 3 * size.Width / 4
	objects[0].Resize(fyne.NewSize(threeFourth, size.Height))
	objects[0].Move(fyne.NewPos(0, 0))

	oneFourth := size.Width - threeFourth
	objects[1].Resize(fyne.NewSize(oneFourth, size.Height))
	objects[1].Move(fyne.NewPos(threeFourth, 0))
}

func (l ThreeFourthOneFourthLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minW := fyne.Max(objects[0].MinSize().Width*4/3, objects[1].MinSize().Width*4)
	minH := fyne.Max(objects[0].MinSize().Height, objects[1].MinSize().Height)
	return fyne.NewSize(minW, minH)
}

// ----- Helper functions -----
func NewOneThirdTwoThird(first, second fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&OneThirdTwoThirdLayout{}, first, second)
}

func NewTwoThirdOneThird(first, second fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&TwoThirdOneThirdLayout{}, first, second)
}

func NewOneFourthThreeFourth(first, second fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&OneFourthThreeFourthLayout{}, first, second)
}

func NewThreeFourthOneFourth(first, second fyne.CanvasObject) fyne.CanvasObject {
	return container.New(&ThreeFourthOneFourthLayout{}, first, second)
}
