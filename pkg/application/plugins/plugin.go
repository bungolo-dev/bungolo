package plugins

type Plugin interface {
	Register()
	Initialize(settings map[string]any)
}
