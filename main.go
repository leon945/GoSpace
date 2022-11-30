package main

import (
	ggo "ggame/game_object"
	"ggame/gtypes"
	"ggame/loop"
	planet_plugins "ggame/planet/plugins"
	"ggame/util"
	"ggame/util/constants"
	"math/rand"
	"time"

	"github.com/gonutz/prototype/draw"
)

var isWindowSet bool = false
var timeKeeper *loop.TimeKeeper

func main() {
	timeKeeper = &loop.TimeKeeper{}
	// test1()
	createRandomStars(1)
	createRandomPlanets(100)
	//createGasGiantsAndMoons(10)
	//createRandomMoons("", gtypes.Vector2_Zero(), 400)
	runGGOStarts()
	timeKeeper.Start()
	draw.RunWindow("Title", constants.ScreenWidth, constants.ScreenHeight, update)
}

func update(window draw.Window) {
	timeKeeper.End()
	timeKeeper.Start()
	if !isWindowSet {
		for _, ggo := range ggo.GameObjects {
			ggo.Window = &window
		}
		isWindowSet = true
	}

	runGGOUpdates()

	// // find the screen center
	// w, h := window.Size()
	// centerX, centerY := w/2, h/2

	// // draw a button in the center of the screen
	// mouseX, mouseY := window.MousePosition()
	// mouseInCircle := math.Hypot(float64(mouseX-centerX), float64(mouseY-centerY)) < 20
	// color := draw.DarkRed
	// if mouseInCircle {
	// 	color = draw.Red
	// }

	// window.FillEllipse(centerX-20, centerY-20, 40, 40, color)
	// window.DrawEllipse(centerX-20, centerY-20, 40, 40, draw.White)
	// if mouseInCircle {
	// 	window.DrawScaledText("Close!", centerX-40, centerY+25, 1.6, draw.Green)
	// }

	// // check all mouse clicks that happened during this frame
	// for _, click := range window.Clicks() {
	// 	dx, dy := click.X-centerX, click.Y-centerY
	// 	squareDist := dx*dx + dy*dy
	// 	if squareDist <= 20*20 {
	// 		// close the window and end the application
	// 		window.Close()
	// 	}
	// }
}

func runGGOStarts() {
	for _, ggo := range ggo.GameObjects {
		ggo.Start()
	}
}

func runGGOUpdates() {
	for _, ggo := range ggo.GameObjects {
		ggo.Update()
	}
}

func createRandomPlanets(amount int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < amount; i++ {
		xPos := rand.Float64() * float64(constants.ScreenWidth)
		yPos := rand.Float64() * float64(constants.ScreenHeight)

		xV := rand.Float64() * 4
		yV := rand.Float64() * 4

		if rand.Intn(100) > 50 {
			xV = -xV
		}

		if rand.Intn(100) > 50 {
			yV = -yV
		}

		var m float64 = constants.MinPlanetMass + (rand.Float64() * (constants.MaxNormalPlanetMass - constants.MinPlanetMass))
		d := planet_plugins.GetDiameterFromMass(constants.PLANET_BODY_TYPE, m)

		r := rand.Float32()
		g := rand.Float32()
		b := rand.Float32()

		nPlanet := planet_plugins.CreatePlanetGameObject(
			float64(xPos),
			float64(yPos),
			d,
			m,
			draw.Color{A: 1, R: r, G: g, B: b},
			gtypes.Vector2{X: float64(xV), Y: float64(yV)},
			timeKeeper,
			"",
		)
		nPPlugin := nPlanet.GetPlugin(util.TypeofObject(&planet_plugins.Planet{})).(*planet_plugins.Planet)
		nPPlugin.SetAsPlanet()
		nPlanet.AllGameObjects = &ggo.GameObjects
		ggo.GameObjects = append(ggo.GameObjects, nPlanet)
	}
}

func createRandomStars(amount int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < amount; i++ {
		xPos := float64(constants.ScreenWidth) / 2  // + rand.Float64()*600
		yPos := float64(constants.ScreenHeight) / 2 // + rand.Float64()*600

		xV := rand.Float64() * .001
		yV := rand.Float64() * 0

		if rand.Intn(100) > 50 {
			xV = -xV
		}

		if rand.Intn(100) > 50 {
			yV = -yV
		}

		var m float64 = constants.MinStarMass + (rand.Float64() * (constants.MaxStarMass - constants.MinStarMass))
		d := planet_plugins.GetDiameterFromMass(constants.STAR_BODY_TYPE, m)

		r := .8 + rand.Float32()*.2
		g := rand.Float32() * .1
		b := rand.Float32() * .5

		nPlanet := planet_plugins.CreatePlanetGameObject(
			float64(xPos),
			float64(yPos),
			d,
			m,
			draw.Color{A: 1, R: r, G: g, B: b},
			gtypes.Vector2{X: float64(xV), Y: float64(yV)},
			timeKeeper,
			"",
		)
		nPPlugin := nPlanet.GetPlugin(util.TypeofObject(&planet_plugins.Planet{})).(*planet_plugins.Planet)
		nPPlugin.SetAsStar()
		nPlanet.AllGameObjects = &ggo.GameObjects
		ggo.GameObjects = append(ggo.GameObjects, nPlanet)
	}
}

func createGasGiantsAndMoons(amount int) {
	for i := 0; i < amount; i++ {
		xPos := rand.Float64() * float64(constants.ScreenWidth)
		yPos := rand.Float64() * float64(constants.ScreenHeight)
		var m float64 = constants.MinGasGiantMass + (rand.Float64() * (constants.MaxGasGiantMass - constants.MinGasGiantMass))
		d := planet_plugins.GetDiameterFromMass(1, m)

		r := rand.Float32()
		g := rand.Float32() * .4
		b := rand.Float32() * .1

		xV := rand.Float64() * 2
		yV := rand.Float64() * .2

		giant := planet_plugins.CreatePlanetGameObject(
			xPos,
			yPos,
			d,
			m,
			draw.Color{A: 1, R: r, G: g, B: b},
			gtypes.Vector2{X: float64(xV), Y: float64(yV)},
			timeKeeper,
			"",
		)
		nPPlugin := giant.GetPlugin(util.TypeofObject(&planet_plugins.Planet{})).(*planet_plugins.Planet)
		nPPlugin.SetAsPlanet()
		giant.AllGameObjects = &ggo.GameObjects
		ggo.GameObjects = append(ggo.GameObjects, giant)

		numOfMoons := 10
		createRandomMoons(nPPlugin.Id, gtypes.Vector2{X: xPos, Y: yPos}, numOfMoons)
	}
}

func createRandomMoons(parentPlanetId string, planetPos gtypes.Vector2, amount int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < amount; i++ {
		xPos := planetPos.X + rand.Float64()*5 + 1
		yPos := planetPos.Y + rand.Float64()*5 + 1

		xV := 1 + rand.Float64()
		yV := rand.Float64() * .1

		if parentPlanetId == "" {
			xPos = rand.Float64() * float64(constants.ScreenWidth)
			yPos = rand.Float64() * float64(constants.ScreenHeight)

			if rand.Intn(100) > 50 {
				xV = -xV
			}

			if rand.Intn(100) > 50 {
				yV = -yV
			}
		}

		var m float64 = constants.MinMoonMass + (rand.Float64() * (constants.MaxMoonMass - constants.MinMoonMass))
		d := planet_plugins.GetDiameterFromMass(constants.MOON_BODY_TYPE, m)

		c := rand.Float32() * .7

		nPlanet := planet_plugins.CreatePlanetGameObject(
			float64(xPos),
			float64(yPos),
			d,
			m,
			draw.Color{A: 1, R: c, G: c, B: c},
			gtypes.Vector2{X: float64(xV), Y: float64(yV)},
			timeKeeper,
			parentPlanetId,
		)
		nPPlugin := nPlanet.GetPlugin(util.TypeofObject(&planet_plugins.Planet{})).(*planet_plugins.Planet)
		nPPlugin.SetAsMoon()
		nPlanet.AllGameObjects = &ggo.GameObjects
		ggo.GameObjects = append(ggo.GameObjects, nPlanet)
	}
}

func test1() {
	createTestPlanet(0, 500, 10, 0, 2000)
	createTestPlanet(800, 500, -10, 0, 2000)
}

func createTestPlanet(xPos float64, yPos float64, xVel float64, yVel float64, mass float64) {
	d := planet_plugins.GetDiameterFromMass(1, mass)

	r := rand.Float32()
	g := rand.Float32() * .4
	b := rand.Float32() * .1

	nPlanet := planet_plugins.CreatePlanetGameObject(
		xPos,
		yPos,
		d,
		mass,
		draw.Color{A: 1, R: r, G: g, B: b},
		gtypes.Vector2{X: float64(xVel), Y: float64(yVel)},
		timeKeeper,
		"",
	)
	nPPlugin := nPlanet.GetPlugin(util.TypeofObject(&planet_plugins.Planet{})).(*planet_plugins.Planet)
	nPPlugin.SetAsPlanet()
	nPlanet.AllGameObjects = &ggo.GameObjects
	ggo.GameObjects = append(ggo.GameObjects, nPlanet)
}
