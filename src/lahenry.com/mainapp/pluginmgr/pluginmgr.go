package pluginmgr

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/hashicorp/go-plugin"
	"lahenry.com/library/baseplugin"
	"lahenry.com/library/baseplugin/greeter"
)

type PluginManager struct{}

func (p *PluginManager) StartPlugin(pluginfile string) error {
	curpwd, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	pluginfile = filepath.Join(curpwd, pluginfile)

	log.Printf("Starting plugin: %s", pluginfile)

	client := plugin.NewClient(&plugin.ClientConfig{
		Cmd:        exec.Command(pluginfile),
		Managed:    true,
		SyncStdout: os.Stdout,
		SyncStderr: os.Stderr,

		HandshakeConfig: baseplugin.HandshakeConfig,
		Plugins:         baseplugin.GetPluginMap(nil),
	})

	rpcclient, err := client.Client()

	if err != nil {
		log.Printf("Failed to get RPC Client: %s", err)
		client.Kill()
		return err
	}

	// get the interface
	raw, err := rpcclient.Dispense(greeter.InterfaceName)
	if err != nil {
		log.Printf("Failed to get interface: %s error: %s", greeter.InterfaceName, err)
		return err
	}

	greeterObj := raw.(greeter.IGreeter)

	go func() {
		for idx := 0; true; idx++ {
			log.Printf("%s - %s - %d - %d", greeterObj.Name(), greeterObj.Version(), greeterObj.StartTime(), idx)
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	return nil
}
