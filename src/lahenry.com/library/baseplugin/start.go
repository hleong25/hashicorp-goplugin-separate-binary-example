package baseplugin

import (
	"log"

	"github.com/hashicorp/go-plugin"
	"lahenry.com/library/baseplugin/greeter"
)

// PluginOpts are options to start the Hashicorp plugin framework
type PluginOpts struct {
	Greeter     greeter.GreeterFunc
	RunAsPlugin bool
}

// StartPlugin starts the Hashicorp go-plugin system over RPC
// between the agent and the plugin, asynchronously.
func StartPlugin(opts *PluginOpts, quit chan bool) {
	if opts.RunAsPlugin {
		go func() {
			log.Println("Starting plugin communication...")

			plugin.Serve(&plugin.ServeConfig{
				HandshakeConfig: HandshakeConfig,
				Plugins:         GetPluginMap(opts),
			})

			log.Println("Exiting plugin communication...")

			quit <- true
			log.Println("Exiting plugin...")
		}()
	} else {
		log.Println("Starting in standalone mode...")
	}
}
