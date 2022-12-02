package constants

const ScreenWidth int = 1200
const ScreenHeight int = 800

var TimeScale float64 = 4000

const DistanceScale float64 = 1000

const MinMoonMass float64 = .001
const MaxMoonMass float64 = .5

const MinPlanetMass float64 = 1
const MaxNormalPlanetMass float64 = 317
const MinGasGiantMass float64 = 200
const MaxGasGiantMass float64 = 400
const MaxPlanetMass float64 = 4200

const MinStarMass float64 = 300000
const MaxStarMass float64 = 600000

const MOON_BODY_TYPE int = 0
const PLANET_BODY_TYPE int = 1
const STAR_BODY_TYPE int = 2

const MaxDistanceForPlanetCollision float64 = 1

func IncreaseTimeScale() {
	if TimeScale < 16000 {
		TimeScale *= 2
	} else {
		TimeScale = 16000
	}
}

func DecreaseTimeScale() {
	if TimeScale > 500 {
		TimeScale /= 2
	} else {
		TimeScale = 500
	}
}
