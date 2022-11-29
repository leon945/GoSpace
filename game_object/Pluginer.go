package gameobject

type Pluginer interface {
	Start()
	Update()
	Destroy()
	SetEnabled(enabled bool)
	IsEnabled() bool
	SetGameObject(ggo *GameObject)
}
