package gplugins

import (
	gameobject "ggame/game_object"
	"ggame/gtypes"
)

type RigidBody struct {
	mass          float64
	gravity       float64
	enabled       bool
	airFriction   float64
	touchFriction float64
	velocity      gtypes.Vector2
	force         gtypes.Vector2
	ggo           *gameobject.GameObject
}

const deltaTimeMultiplier float64 = 10
const distanceScale float64 = 1000

func (rb *RigidBody) Init(gameObject *gameobject.GameObject, mass float64, airFriction float64, touchFriction float64) {
	rb.SetGameObject(gameObject)
	rb.SetEnabled(true)
	rb.SetAirFriction(airFriction)
	rb.SetTouchFriction(touchFriction)
	rb.SetMass(mass)
}

func (rb *RigidBody) Start() {

}

func (rb *RigidBody) Update() {
	rb.applyForce()
	rb.applyGravity()
	rb.applyVelocity()
}

func (rb *RigidBody) applyForce() {
	yChange := rb.force.Y * rb.ggo.TimeKeeper.DeltaTime() * deltaTimeMultiplier * (1 / rb.mass) * distanceScale
	xChange := rb.force.X * rb.ggo.TimeKeeper.DeltaTime() * deltaTimeMultiplier * (1 / rb.mass) * distanceScale
	rb.velocity = gtypes.Vector2{X: rb.velocity.X + xChange, Y: rb.velocity.Y + yChange}
}

func (rb *RigidBody) applyGravity() {
	yChange := rb.gravity * rb.ggo.TimeKeeper.DeltaTime() * deltaTimeMultiplier * distanceScale
	rb.velocity = gtypes.Vector2{X: rb.velocity.X, Y: rb.velocity.Y + yChange}
}

func (rb *RigidBody) applyVelocity() {
	if rb.velocity.Magnitude() > 0 {
		cp := rb.ggo.Position
		cp.Val.X += rb.velocity.X * rb.ggo.TimeKeeper.DeltaTime() * deltaTimeMultiplier * (1 / distanceScale)
		cp.Val.Y += rb.velocity.Y * rb.ggo.TimeKeeper.DeltaTime() * deltaTimeMultiplier * (1 / distanceScale)
	}
}

func (rb *RigidBody) Destroy() {

}

func (rb *RigidBody) SetEnabled(enabled bool) {
	rb.enabled = enabled
}

func (rb *RigidBody) IsEnabled() bool {
	return rb.enabled
}

func (rb *RigidBody) SetGameObject(ggo *gameobject.GameObject) {
	rb.ggo = ggo
}

func (rb *RigidBody) SetGravity(gravity float64) {
	rb.gravity = gravity
}

func (rb *RigidBody) SetMass(mass float64) {
	rb.mass = mass
}

func (rb *RigidBody) GetMass() float64 {
	return rb.mass
}

func (rb *RigidBody) GetVelocity() gtypes.Vector2 {
	return rb.velocity
}

func (rb *RigidBody) GetForce() gtypes.Vector2 {
	return rb.force
}

func (rb *RigidBody) SetAirFriction(airFriction float64) {
	rb.airFriction = airFriction
}

func (rb *RigidBody) SetTouchFriction(touchFriction float64) {
	rb.touchFriction = touchFriction
}

func (rb *RigidBody) SetVelocity(velocity gtypes.Vector2) {
	rb.velocity = velocity
}

func (rb *RigidBody) SetForce(force gtypes.Vector2) {
	rb.force = force
}
