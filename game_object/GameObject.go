package gameobject

import (
	"ggame/gtypes"
	"ggame/loop"
	"ggame/util"

	"github.com/gonutz/prototype/draw"
	"github.com/google/uuid"
)

type GameObject struct {
	Position       *Position
	Size           *Size
	Enabled        bool
	plugins        []Pluginer
	Window         *draw.Window
	TimeKeeper     *loop.TimeKeeper
	AllGameObjects *[]*GameObject
	Id             string
}

var GameObjects []*GameObject
var GameObjectsToDelete []*GameObject

func (g *GameObject) Init(xPos float64, yPos float64, width float64, height float64, enabled bool, timeKeeper *loop.TimeKeeper) {
	g.Position = &Position{Val: gtypes.Vector2{X: xPos, Y: yPos}}
	g.Size = &Size{Val: gtypes.Vector2{X: width, Y: height}}
	g.Enabled = enabled
	g.TimeKeeper = timeKeeper
	g.Id = uuid.New().String()
}

func (g *GameObject) AddPlugin(plugin Pluginer) {
	g.plugins = append(g.plugins, plugin)
}

func (g *GameObject) Intersects(other *GameObject) bool {
	if (g.GetTopLeft().X < other.GetBottomRight().X) &&
		(g.GetBottomRight().X > other.GetTopLeft().X) &&
		(g.GetTopLeft().Y < other.GetBottomRight().Y) &&
		(g.GetBottomRight().Y > other.GetTopLeft().Y) {
		return true
	}

	return false
}

func (g *GameObject) GetTopLeft() gtypes.Vector2 {
	tlX := g.Position.Val.X - (g.Size.Val.X / 2)
	tlY := g.Position.Val.Y - (g.Size.Val.Y / 2)

	return gtypes.Vector2{X: tlX, Y: tlY}
}

func (g *GameObject) GetTopRight() gtypes.Vector2 {
	tlX := g.Position.Val.X + (g.Size.Val.X / 2)
	tlY := g.Position.Val.Y - (g.Size.Val.Y / 2)

	return gtypes.Vector2{X: tlX, Y: tlY}
}

func (g *GameObject) GetBottomLeft() gtypes.Vector2 {
	tlX := g.Position.Val.X - (g.Size.Val.X / 2)
	tlY := g.Position.Val.Y + (g.Size.Val.Y / 2)

	return gtypes.Vector2{X: tlX, Y: tlY}
}

func (g *GameObject) GetBottomRight() gtypes.Vector2 {
	tlX := g.Position.Val.X + (g.Size.Val.X / 2)
	tlY := g.Position.Val.Y + (g.Size.Val.Y / 2)

	return gtypes.Vector2{X: tlX, Y: tlY}
}

func (g *GameObject) Start() {
	if g.Enabled {
		for _, el := range g.plugins {
			if el != nil {
				(el).Start()
			}
		}
	}
}

func (g *GameObject) Update() {
	if g != nil && g.Enabled && g.Window != nil {
		for _, el := range g.plugins {
			if el != nil {
				(el).Update()
			}
		}
	}

	deletedMarkedGameObjects()
}

func (g *GameObject) Destroy() {
	for _, el := range g.plugins {
		if el != nil {
			(el).Destroy()
		}
	}
	g.Enabled = false
}

func (g *GameObject) GetPlugin(pluginType string) Pluginer {
	for _, el := range g.plugins {
		if util.TypeofObject(el) == pluginType {
			return el
		}
	}

	return nil
}

func (g *GameObject) GetAllPlugins(pluginType string) []Pluginer {
	var plugins []Pluginer
	for _, el := range g.plugins {
		if util.TypeofObject(el) == pluginType {
			plugins = append(plugins, el)
		}
	}

	return plugins
}

func FindAllPlugins(pluginType string) []Pluginer {
	var plugins []Pluginer
	for _, el := range GameObjects {
		goPlugins := el.GetAllPlugins(pluginType)
		if len(goPlugins) > 0 {
			plugins = append(plugins, goPlugins...)
		}
	}

	return plugins
}

func MarkForDeletion(ggo *GameObject) {
	GameObjectsToDelete = append(GameObjectsToDelete, ggo)
}

func deletedMarkedGameObjects() {
	for _, ggoToDelete := range GameObjectsToDelete {
		var i int = -1

		for index, gameObject := range GameObjects {
			if gameObject.Id == ggoToDelete.Id {
				i = index
				ggoToDelete.Destroy()
				break
			}
		}

		if i < 0 {
			continue
		}

		// Remove the element at index i from a.
		GameObjects[i] = GameObjects[len(GameObjects)-1] // Copy last element to index i.
		GameObjects[len(GameObjects)-1] = nil            // Erase last element (write zero value).
		GameObjects = GameObjects[:len(GameObjects)-1]   // Truncate slice.
	}
}
