package planet

import (
	ggo "ggame/game_object"
	"ggame/gtypes"
)

type Planeter interface {
	IsPlanet() bool
	GetBodyType() int
	GetGameObject() *ggo.GameObject
	GetMass() float64
	GetPosition() gtypes.Vector2
}
