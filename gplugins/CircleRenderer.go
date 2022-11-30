package gplugins

import (
	gameobject "ggame/game_object"

	"github.com/gonutz/prototype/draw"
)

type CircleRenderer struct {
	Renderer
	ggo     *gameobject.GameObject
	filled  bool
	enabled bool
}

func (cr *CircleRenderer) Init(gameObject *gameobject.GameObject, filled bool, color draw.Color) {
	cr.SetGameObject(gameObject)
	cr.SetEnabled(true)
	cr.SetColor(color)
	cr.SetFilled(filled)
}

func (cr *CircleRenderer) Start() {

}

func (cr *CircleRenderer) Update() {
	if !cr.enabled {
		return
	}

	topLeft := cr.ggo.GetTopLeft()
	size := cr.ggo.Size.Val

	if cr.filled {
		(*cr.ggo.Window).FillEllipse(int(topLeft.X), int(topLeft.Y), int(size.X), int(size.Y), cr.color)
	}

	(*cr.ggo.Window).DrawEllipse(int(topLeft.X), int(topLeft.Y), int(size.X), int(size.Y), cr.color)
}

func (cr *CircleRenderer) Destroy() {

}

func (cr *CircleRenderer) SetEnabled(enabled bool) {
	cr.enabled = enabled
}

func (cr *CircleRenderer) IsEnabled() bool {
	return cr.enabled
}

func (cr *CircleRenderer) SetGameObject(ggo *gameobject.GameObject) {
	cr.ggo = ggo
}

func (cr *CircleRenderer) SetColor(color draw.Color) {
	cr.color = color
}

func (cr *CircleRenderer) SetFilled(filled bool) {
	cr.filled = filled
}
