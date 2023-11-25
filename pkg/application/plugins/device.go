package plugins

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type Device interface {
	Initialize(settings map[string]string) InitializeResult
	Query() string
	Kill()
}

type InitializeResult int

// Bungolow local implmentation of a device plugin.
// The bungolow framework will registered the device and initialize the device.
// Once initialized the device features can be invoked.
type DeviceRPC struct{ client *rpc.Client }

func (g *DeviceRPC) Query() string {

	var resp string
	err := g.client.Call("Plugin.Query", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}

	return resp
}

func (g *DeviceRPC) Initialize(settings map[string]string) InitializeResult {

	var resp InitializeResult
	err := g.client.Call("Plugin.Initialize", settings, &resp)
	if err != nil {
		return 0
	}
	return resp
}

func (g *DeviceRPC) Kill() {

	g.client.Call("Plugin.Kil", new(interface{}), nil)
}

// Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type DeviceRPCServer struct {
	// This is the real implementation
	Impl Device
}

func (s *DeviceRPCServer) Query(args interface{}, resp *string) error {
	*resp = s.Impl.Query()
	return nil
}

func (s *DeviceRPCServer) Initialize(args map[string]string, resp *InitializeResult) error {
	*resp = s.Impl.Initialize(args)
	return nil
}

func (s *DeviceRPCServer) Kill(args interface{}, resp *string) error {
	s.Impl.Kill()
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a DeviceRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return DeviceRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type DevicePlugin struct {
	// Impl Injection
	Impl Device
}

func (p *DevicePlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &DeviceRPCServer{Impl: p.Impl}, nil
}

func (DevicePlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &DeviceRPC{client: c}, nil
}
