package device

type Device interface {
	DeviceId()
	Query() Capabilities
}

type Capabilities struct {
}
