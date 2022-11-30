package gplugins_PlanetRenderer

import (
	gameobject "ggame/game_object"
	"ggame/gplugins"

	"github.com/gonutz/prototype/draw"
)

type PlanetRenderer struct {
	gplugins.Renderer
	ggo     *gameobject.GameObject
	enabled bool
	color   draw.Color
	star    bool
}

func (cr *PlanetRenderer) Init(ggo *gameobject.GameObject, color draw.Color, isStar bool) {
	cr.SetGameObject(ggo)
	cr.SetEnabled(true)
	cr.SetColor(color)
	cr.star = isStar
}

func (cr *PlanetRenderer) Start() {

}

func (cr *PlanetRenderer) Update() {
	if !cr.enabled {
		return
	}

	ggo := cr.ggo

	topLeft := ggo.GetTopLeft()
	size := ggo.Size.Val

	(*ggo.Window).FillEllipse(int(topLeft.X), int(topLeft.Y), int(size.X), int(size.Y), cr.color)
	(*ggo.Window).DrawEllipse(int(topLeft.X), int(topLeft.Y), int(size.X), int(size.Y), cr.color)

	if !cr.star {
		return
	}

	l1Xf := ggo.Position.Val.X
	l1Yf := ggo.Position.Val.Y - ggo.Size.Val.Y/2
	l1Xt := l1Xf
	l1Yt := l1Yf - ggo.Size.Val.Y/3
	(*ggo.Window).DrawLine(int(l1Xf), int(l1Yf), int(l1Xt), int(l1Yt), cr.color)

	l2Xf := ggo.Position.Val.X
	l2Yf := ggo.Position.Val.Y + ggo.Size.Val.Y/2
	l2Xt := l2Xf
	l2Yt := l2Yf + ggo.Size.Val.Y/3
	(*ggo.Window).DrawLine(int(l2Xf), int(l2Yf), int(l2Xt), int(l2Yt), cr.color)

	l3Xf := ggo.Position.Val.X - ggo.Size.Val.X/2
	l3Yf := ggo.Position.Val.Y
	l3Xt := l3Xf - ggo.Size.Val.X/3
	l3Yt := l3Yf
	(*ggo.Window).DrawLine(int(l3Xf), int(l3Yf), int(l3Xt), int(l3Yt), cr.color)

	l4Xf := ggo.Position.Val.X + ggo.Size.Val.X/2
	l4Yf := ggo.Position.Val.Y
	l4Xt := l4Xf + ggo.Size.Val.X/3
	l4Yt := l4Yf
	(*ggo.Window).DrawLine(int(l4Xf), int(l4Yf), int(l4Xt), int(l4Yt), cr.color)
}

func (cr *PlanetRenderer) Destroy() {

}

func (cr *PlanetRenderer) SetEnabled(enabled bool) {
	cr.enabled = enabled
}

func (cr *PlanetRenderer) IsEnabled() bool {
	return cr.enabled
}

func (cr *PlanetRenderer) SetGameObject(ggo *gameobject.GameObject) {
	cr.ggo = ggo
}

func (cr *PlanetRenderer) SetColor(color draw.Color) {
	cr.color = color
}
