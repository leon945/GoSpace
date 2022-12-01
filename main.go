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
var mouseDownPos gtypes.Vector2
var mouseUpPos gtypes.Vector2
var startMouseBtn draw.MouseButton
var isMouseDown bool = false

func main() {
	timeKeeper = &loop.TimeKeeper{}
	// test1()
	//createRandomStars(1)
	//createRandomPlanets(200)
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
	handleMouse(window)
	handleKeyboard(window)
}

func handleMouse(window draw.Window) {
	// check all mouse clicks that happened during this frame
	rand.Seed(time.Now().UnixNano())
	if !isMouseDown {
		clicks := window.Clicks()
		if len(clicks) > 0 {
			click := window.Clicks()[0]
			if click.Button == draw.MiddleButton {
				createRandomPlanets(100, &window)
			} else {
				isMouseDown = true
				startMouseBtn = click.Button
				mouseDownPos = gtypes.Vector2{X: float64(click.X), Y: float64(click.Y)}
			}
		}
	}

	if isMouseDown {
		x, y := window.MousePosition()
		if startMouseBtn == draw.LeftButton {
			window.DrawLine(int(mouseDownPos.X), int(mouseDownPos.Y), x, y, draw.White)
		} else {
			window.DrawLine(int(mouseDownPos.X), int(mouseDownPos.Y), x, y, draw.Red)
		}
	}

	if isMouseDown && !window.IsMouseDown(startMouseBtn) {
		isMouseDown = false
		x, y := window.MousePosition()
		mouseUpPos = gtypes.Vector2{X: float64(x), Y: float64(y)}

		mass := constants.MinPlanetMass + (rand.Float64() * (constants.MaxNormalPlanetMass - constants.MinPlanetMass))
		if startMouseBtn == draw.RightButton {
			mass = constants.MinStarMass + (rand.Float64() * (constants.MaxStarMass - constants.MinStarMass))
		}

		xVel := (mouseUpPos.X - mouseDownPos.X) * .01
		yVel := (mouseUpPos.Y - mouseDownPos.Y) * .01

		createTestPlanet(float64(mouseDownPos.X), float64(mouseDownPos.Y), xVel, yVel, mass, &window)
	}
}

func handleKeyboard(window draw.Window) {
	if window.IsKeyDown(draw.KeyEscape) {
		for _, gameObj := range ggo.GameObjects {
			ggo.MarkForDeletion(gameObj)
		}
	}

	if window.IsKeyDown(draw.KeyQ) {
		window.Close()
	}
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

func createRandomPlanets(amount int, window *draw.Window) {
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

		color := planet_plugins.GetColorFromBodyType(constants.PLANET_BODY_TYPE)

		nPlanet := planet_plugins.CreatePlanetGameObject(
			float64(xPos),
			float64(yPos),
			d,
			m,
			color,
			gtypes.Vector2{X: float64(xV), Y: float64(yV)},
			timeKeeper,
			"",
		)

		nPlanet.Window = window
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

		color := planet_plugins.GetColorFromBodyType(constants.STAR_BODY_TYPE)

		nPlanet := planet_plugins.CreatePlanetGameObject(
			float64(xPos),
			float64(yPos),
			d,
			m,
			color,
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

		color := planet_plugins.GetColorFromBodyType(constants.PLANET_BODY_TYPE)

		xV := rand.Float64() * 2
		yV := rand.Float64() * .2

		giant := planet_plugins.CreatePlanetGameObject(
			xPos,
			yPos,
			d,
			m,
			color,
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

		color := planet_plugins.GetColorFromBodyType(constants.MOON_BODY_TYPE)

		nPlanet := planet_plugins.CreatePlanetGameObject(
			float64(xPos),
			float64(yPos),
			d,
			m,
			color,
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
	createTestPlanet(0, 500, 10, 0, 2000, nil)
	createTestPlanet(800, 500, -10, 0, 2000, nil)
}

func createTestPlanet(xPos float64, yPos float64, xVel float64, yVel float64, mass float64, window *draw.Window) {
	var bodyType int = constants.PLANET_BODY_TYPE
	if mass > constants.MinStarMass {
		bodyType = constants.STAR_BODY_TYPE
	}
	d := planet_plugins.GetDiameterFromMass(bodyType, mass)

	color := planet_plugins.GetColorFromBodyType(bodyType)

	nPlanet := planet_plugins.CreatePlanetGameObject(
		xPos,
		yPos,
		d,
		mass,
		color,
		gtypes.Vector2{X: float64(xVel), Y: float64(yVel)},
		timeKeeper,
		"",
	)
	if window != nil {
		nPlanet.Window = window
	}

	nPPlugin := nPlanet.GetPlugin(util.TypeofObject(&planet_plugins.Planet{})).(*planet_plugins.Planet)
	nPPlugin.SetAsPlanet()
	nPlanet.AllGameObjects = &ggo.GameObjects
	ggo.GameObjects = append(ggo.GameObjects, nPlanet)
}
