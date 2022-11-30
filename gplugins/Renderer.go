package gplugins

import "github.com/gonutz/prototype/draw"

type Renderer struct {
	color draw.Color
}

func (r *Renderer) SetColor(color *draw.Color) {
	r.color = *color
}
