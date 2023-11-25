package main

import (
	"os"
	"strconv"

	"github.com/bungolow-dev/bungolow/pkg/application/plugins"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Roku struct {
	ip     string
	port   int
	logger hclog.Logger
}

func (r *Roku) Initialize(settings map[string]string) plugins.InitializeResult {

	p, err := strconv.Atoi(settings["port"])

	if err != nil {
		roku.ip = string(settings["ip"])
		roku.port = 8060
		r.logger.Debug("Configured Roku settings %+v", settings)
		return 0
	}

	roku.ip = string(settings["ip"])
	roku.port = p
	r.logger.Debug("Configured Roku settings %+v", settings)

	return 1
}

func (r *Roku) Query() string {
	r.logger.Debug("Attempting to Query Device Info")

	pressHome(r.ip)
	q := queryInfo(r.ip)
	return "DEVICE INFORMATION: \n\n\n" + q
}

func (r *Roku) Kill() {
	r.logger.Debug("Attempting to Query Device Info")
	KillClient()
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

var roku *Roku

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	roku = &Roku{
		logger: logger,
	}

	var pluginMap = map[string]plugin.Plugin{
		"device": &plugins.DevicePlugin{Impl: roku},
	}

	logger.Debug("Bungolow Loaded Roku Plugin")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
