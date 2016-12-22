package greeter

import (
	"log"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

type greeterClient struct {
	Broker *plugin.MuxBroker
	Client *rpc.Client
}

// Client implmentation of go-plugin.plugin.Plugin.Client
func (p *Greeter) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &greeterClient{Broker: b, Client: c}, nil
}

////////////////////////////////////////////////////////////////////////////////
//
//  For each function defined in the interface, there needs to be an equivalent
//  greeterClient call
//
////////////////////////////////////////////////////////////////////////////////

// Name initiates an RPC call to the plugin name
func (p *greeterClient) Name() string {
	var resp string
	err := p.Client.Call("Plugin.Name", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
		// TODO: log.Fatal() will exit the process automatically.
		// Need to figure out what the proper thing to do is
	}
	// log.Println("goplugin-client.Name():", resp)
	return resp
}

// Version initiates an RPC call to the plugin version
func (p *greeterClient) Version() string {
	var resp string
	err := p.Client.Call("Plugin.Version", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
		// TODO: log.Fatal() will exit the process automatically.
		// Need to figure out what the proper thing to do is
	}
	// log.Println("goplugin-client.Version():", resp)
	return resp
}

// StartTime initiates an RPC call to the plugin start time
func (p *greeterClient) StartTime() int64 {
	var resp int64
	err := p.Client.Call("Plugin.StartTime", new(interface{}), &resp)
	if err != nil {
		log.Fatal(err)
		// TODO: log.Fatal() will exit the process automatically.
		// Need to figure out what the proper thing to do is
	}
	// log.Println("goplugin-client.StartTime():", resp)
	return resp
}
