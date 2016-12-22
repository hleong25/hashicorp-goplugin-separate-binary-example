package baseplugin

// Configuration values for Hashicorp go-plugin framework

import (
	"github.com/hashicorp/go-plugin"
	"lahenry.com/library/baseplugin/greeter"
)

// HandshakeConfig is used to perform a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "HASHICORP_GOPLUGIN_SEPARAGE_BINARY_EXAMPLE",
	MagicCookieValue: "6219b2ed05eaa8a4063ecde963ad211fecd58afecbbbf2a3225ac54ec156eb0c", // sha256('henry leong')
}

// GetPluginMap returns the plugin map defined Hashicorp go-plugin.
// The reserved parameter should only be used by the RPC receiver (the plugin).
// Otherwise, reserved should be nil for the RPC sender (the mainapp).
func GetPluginMap(reserved *PluginOpts) map[string]plugin.Plugin {
	var greeterObj greeter.Greeter

	if reserved != nil {
		greeterObj.F = reserved.Greeter
	}

	return map[string]plugin.Plugin{
		greeter.InterfaceName: &greeterObj,
	}
}
