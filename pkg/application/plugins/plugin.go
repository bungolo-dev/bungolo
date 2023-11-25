package plugins

type Plugin interface {
	Register() error
	Initialize(settings map[string]interface{})
}
