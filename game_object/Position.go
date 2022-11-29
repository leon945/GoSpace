package gameobject

import (
	"fmt"
	"ggame/gtypes"
)

type Position struct {
	Val gtypes.Vector2
}

func (pos Position) Print() {
	fmt.Printf(`Mouse Position is: %g, %g`, pos.Val.X, pos.Val.Y)
}
