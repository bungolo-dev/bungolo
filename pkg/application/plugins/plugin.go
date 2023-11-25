package plugins

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

// Move to database, load configured plugins from db
var pluginCfg map[string]PluginConfig = map[string]PluginConfig{
	"roku1": {Id: "roku1", ExePath: "../roku/roku.exe", Type: "device", Settings: map[string]string{"ip": "10.0.0.62", "port": "8060"}},
	"roku2": {Id: "roku2", ExePath: "../roku/roku.exe", Type: "device", Settings: map[string]string{"ip": "10.0.0.245", "port": "8060"}},
}

var devices map[string]*DeviceDriver = map[string]*DeviceDriver{}

func StartPlugins() error {

	for k, p := range pluginCfg {

		logger := hclog.New(&hclog.LoggerOptions{
			Name:   "PLUGIN: " + k,
			Output: os.Stdout,
			Level:  hclog.Debug,
		})

		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: handshakeConfig,
			Plugins:         pluginMap,
			Cmd:             exec.Command(p.ExePath),
			Logger:          logger,
		})

		rpcClient, err := client.Client()
		if err != nil {
			log.Fatal(err)
			return err
		}

		raw, err := rpcClient.Dispense(p.Type)
		if err != nil {
			log.Fatal(err)
			return err
		}
		device := raw.(Device)
		device.Initialize(p.Settings)
		devices[k] = &DeviceDriver{device: raw.(Device), client: client}
	}

	return nil
}

func LoadPlugin(pluginConf PluginConfig) error {

	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "PLUGIN: " + pluginConf.Id,
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(pluginConf.ExePath),
		Logger:          logger,
	})

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
		return err
	}

	raw, err := rpcClient.Dispense(pluginConf.Type)
	if err != nil {
		log.Fatal(err)
		return err
	}
	device := raw.(Device)
	device.Initialize(pluginConf.Settings)
	devices[pluginConf.Id] = &DeviceDriver{device: raw.(Device), client: client}

	return nil
}

func KillPlugin(id string) error {

	plugin := devices[id]
	if plugin == nil {
		return fmt.Errorf("PLUGIN %s NOT FOUND IN LOADED DEVICES", id)
	}

	delete(devices, id)
	plugin.client.Kill()
	return nil
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

// pluginMap is the map of plugins we can dispense.
var pluginMap = map[string]plugin.Plugin{
	"device": &DevicePlugin{},
}

type PluginConfig struct {
	Id       string            `json:"id"`
	ExePath  string            `json:"exePath"`
	Type     string            `json:"type"`
	Settings map[string]string `json:"settings"`
}

type DeviceDriver struct {
	client *plugin.Client
	device Device
}
