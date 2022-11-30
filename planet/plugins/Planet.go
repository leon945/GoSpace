package planet_plugins

import (
	gameobject "ggame/game_object"
	"ggame/gplugins"
	gplugins_PlanetRenderer "ggame/gplugins/PlanetRenderer"
	"ggame/gtypes"
	"ggame/loop"
	"ggame/util"
	"math"
	"math/rand"
	"time"

	"github.com/gonutz/prototype/draw"
)

type Planet struct {
	enabled        bool
	ggo            *gameobject.GameObject
	rigidBody      *gplugins.RigidBody
	Id             string
	bodyType       int
	ParentPlanetId string
}

func (p *Planet) IsPlanet() bool {
	return true
}

func (p *Planet) Init() {
	p.Id = p.ggo.Id
	rb := (*p.ggo).GetPlugin(util.TypeofObject(&gplugins.RigidBody{}))
	if val, ok := rb.(*gplugins.RigidBody); ok {
		p.rigidBody = val
	}
}

func (p *Planet) Start() {

}

func (p *Planet) Update() {
	if !p.enabled {
		return
	}

	allPlanets := gameobject.FindAllPlugins(util.TypeofObject(&Planet{}))
	p.applyGravityForce(&allPlanets)
	p.checkForCollisions(&allPlanets)
	p.checkIfEscapedSystem()
}

func (p *Planet) checkIfEscapedSystem() {
	x := math.Abs(p.GetPosition().X)
	y := math.Abs(p.GetPosition().Y)

	if x > 10000 || y > 10000 {
		gameobject.MarkForDeletion(p.ggo)
	}
}

func (p *Planet) checkForCollisions(allPlanetsPlugins *[]gameobject.Pluginer) {
	var allPlanets []*Planet
	var collidingPlanets []*Planet
	var collidedMass float64

	var collidedForces gtypes.Vector2 = gtypes.Vector2_Zero()
	var newXPos float64 = 0
	var newYPos float64 = 0
	for _, planetPlugin := range *allPlanetsPlugins {
		planet := planetPlugin.(*Planet)
		allPlanets = append(allPlanets, planet)
		if planet.Id != p.Id {
			vectorToPlanet := p.getVectorTo(planet)
			distanceToPlanet := vectorToPlanet.Magnitude()

			if distanceToPlanet < 1 {
				collidingPlanets = append(collidingPlanets, planet)
				collidedMass += planet.GetMass()
				newXPos += planet.GetPosition().X
				newYPos += planet.GetPosition().Y
			}
		}
	}

	if len(collidingPlanets) == 0 {
		return
	}

	newTotalMass := collidedMass + p.GetMass()

	for _, cp := range collidingPlanets {
		collidedVel := cp.GetVelocity()
		collidedForce := collidedVel.ChangeMagnitudeBy(cp.GetMass() / newTotalMass)
		collidedForces = *collidedForces.Add(*collidedForce)
	}

	newXPos += p.GetPosition().X
	newYPos += p.GetPosition().Y

	newXPos /= float64(len(collidingPlanets) + 1)
	newYPos /= float64(len(collidingPlanets) + 1)

	newTotalForce := p.rigidBody.GetForce()
	newTotalForce = *newTotalForce.Add(collidedForces)

	newVelocity := p.GetVelocity()
	newVelocity = *newVelocity.ChangeMagnitudeBy((p.GetMass() / newTotalMass))
	newVelocity = *newVelocity.Add(collidedForces)

	rand.Seed(time.Now().UnixNano())
	r := rand.Float32()
	g := rand.Float32()
	b := rand.Float32()
	d := GetDiameterFromMass(1, newTotalMass)

	if newTotalMass < 1 {
		r = r * .7
		g = r
		b = r
		d = GetDiameterFromMass(0, newTotalMass)
	} else if newTotalMass > 4200 {
		r = .8 + rand.Float32()*.2
		g = rand.Float32() * .1
		b = rand.Float32() * .5
		d = GetDiameterFromMass(2, newTotalMass)
	}

	nPlanet := CreatePlanetGameObject(
		newXPos,
		newYPos,
		d,
		newTotalMass,
		draw.Color{A: 1, R: r, G: g, B: b},
		newVelocity,
		p.ggo.TimeKeeper,
		p.ParentPlanetId,
	)
	nPPlugin := nPlanet.GetPlugin(util.TypeofObject(&Planet{})).(*Planet)
	if newTotalMass < 1 {
		nPPlugin.SetAsMoon()
	} else if newTotalMass > 1 && newTotalMass < 5000 {
		nPPlugin.SetAsPlanet()
	} else {
		nPPlugin.SetAsStar()
	}

	nPPlugin.SetAsPlanet()
	nPPlugin.rigidBody.SetForce(newTotalForce)

	var planetsToKill []*Planet
	planetsToKill = append(planetsToKill, collidingPlanets...)
	planetsToKill = append(planetsToKill, p)
	for _, planetToKill := range planetsToKill {
		gameobject.MarkForDeletion(planetToKill.ggo)
	}
	UpdatePlanetParents(allPlanets, planetsToKill, nPlanet.Id)

	nPlanet.AllGameObjects = &gameobject.GameObjects
	nPlanet.Window = p.ggo.Window
	gameobject.GameObjects = append(gameobject.GameObjects, nPlanet)
}

func (p *Planet) applyGravityForce(allPlanets *[]gameobject.Pluginer) {
	var gravForces []gtypes.Vector2
	for _, planetPlugin := range *allPlanets {
		planet := planetPlugin.(*Planet)
		if planet.Id != p.Id {
			pForce := p.getGravityForce(planet)
			gravForces = append(gravForces, pForce)
		}
	}

	var resultingForce gtypes.Vector2 = gtypes.Vector2_Zero()
	for _, fv2 := range gravForces {
		resultingForce = *resultingForce.Add(fv2)
	}

	p.rigidBody.SetForce(resultingForce)
}

func (p *Planet) getGravityForce(other *Planet) gtypes.Vector2 {
	otherMass := other.GetMass()

	centerGravity := p.getVectorTo(other)
	distanceFromCenter := centerGravity.Magnitude()
	centerGravityN := *centerGravity.Normalize()
	scaledDistance := distanceFromCenter * p.getDistanceScale(other)

	gravityValue := (1 / math.Pow(scaledDistance, 2)) * otherMass

	if gravityValue <= 0 {
		gravityValue = 0
	}

	centerGravityF := *centerGravityN.ChangeMagnitudeBy(gravityValue)
	return centerGravityF
}

func (p *Planet) getVectorTo(other *Planet) gtypes.Vector2 {
	otherPos := other.GetPosition()

	oPos := p.ggo.Position.Val
	cgX := -oPos.X + otherPos.X
	cgY := -oPos.Y + otherPos.Y
	otherVector := gtypes.Vector2{X: cgX, Y: cgY}
	return otherVector
}

func (p *Planet) getDistanceScale(other *Planet) float64 {
	var scale float64 = 1000
	// oBodyType := other.bodyType
	// switch p.bodyType {
	// case 0: //MOONS
	// 	if other.Id == p.ParentPlanetId {
	// 		scale = 40
	// 	} else {
	// 		scale = 1000000
	// 	}
	// case 1: //PLANETS
	// 	switch oBodyType {
	// 	case 0: //MOONS
	// 		scale = 1
	// 	case 1: //PLANETS
	// 		scale = 1
	// 	case 2: //STARS
	// 		scale = 1
	// 	}
	// case 2: //STARS
	// 	switch oBodyType {
	// 	case 0: //MOONS
	// 		scale = 1
	// 	case 1: //PLANETS
	// 		scale = 1
	// 	case 2: //STARS
	// 		scale = 1
	// 	}
	// }

	return scale
}

func (p *Planet) Destroy() {
	p.enabled = false
}

func (p *Planet) SetEnabled(enabled bool) {
	p.enabled = enabled
}

func (p *Planet) IsEnabled() bool {
	return p.enabled
}

func (p *Planet) SetGameObject(ggo *gameobject.GameObject) {
	p.ggo = ggo
}

func (p *Planet) GetGameObject() *gameobject.GameObject {
	return p.ggo
}

func (p *Planet) GetMass() float64 {
	return (*p.rigidBody).GetMass()
}

func (p *Planet) GetPosition() gtypes.Vector2 {
	return p.ggo.Position.Val
}

func (p *Planet) GetVelocity() gtypes.Vector2 {
	return (*p.rigidBody).GetVelocity()
}

func (p *Planet) SetAsMoon() {
	p.bodyType = 0
}

func (p *Planet) SetAsPlanet() {
	p.bodyType = 1
}

func (p *Planet) SetAsStar() {
	p.bodyType = 2
}

func CreatePlanetGameObject(
	xPos float64,
	yPos float64,
	diameter float64,
	mass float64,
	color draw.Color,
	initialVelocity gtypes.Vector2,
	timeKeeper *loop.TimeKeeper,
	parentPlanetId string,
) *gameobject.GameObject {
	ggo := gameobject.GameObject{}
	ggo.Init(xPos, yPos, diameter, diameter, true, timeKeeper)

	rigidBody := gplugins.RigidBody{}
	rigidBody.Init(&ggo, mass, .5, 1)
	rigidBody.SetGravity(0)
	rigidBody.SetVelocity(initialVelocity)
	ggo.AddPlugin(&rigidBody)

	planet := Planet{
		ParentPlanetId: parentPlanetId,
	}
	planet.SetEnabled(true)
	planet.SetGameObject(&ggo)
	planet.Init()
	ggo.AddPlugin(&planet)

	renderer := &gplugins_PlanetRenderer.PlanetRenderer{}
	renderer.Init(&ggo, color, mass > 4200)
	ggo.AddPlugin(renderer)

	return &ggo
}

func UpdatePlanetParents(allPlanets []*Planet, dyingPlanets []*Planet, idOfNewParent string) {
	for _, p := range allPlanets {
		for _, dp := range dyingPlanets {
			if p.Id != dp.Id && p.ParentPlanetId == dp.Id {
				p.ParentPlanetId = idOfNewParent
			}
		}
	}
}

func GetDiameterFromMass(bodyType int, mass float64) float64 {
	var d float64 = 1
	switch bodyType {
	case 0: //MOONS
		d = 5 + math.Sqrt(mass/math.Pi)*2
	case 1: //PLANETS
		d = 5 + math.Sqrt((mass)/math.Pi)*2
	case 2: //STARS
		d = 10 + math.Sqrt((mass*.02)/math.Pi)*2
		d *= .4
	}

	//d *= .4

	return d
}
