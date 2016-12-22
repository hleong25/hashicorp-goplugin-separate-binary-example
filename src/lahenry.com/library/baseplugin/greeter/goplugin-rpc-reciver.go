package greeter

import (
	"errors"

	"github.com/hashicorp/go-plugin"
)

type greeterServer struct {
	Broker   *plugin.MuxBroker
	IGreeter IGreeter
}

// Server implmentation of go-plugin.plugin.Plugin.Server
func (p *Greeter) Server(b *plugin.MuxBroker) (interface{}, error) {
	if p.F == nil {
		return nil, errors.New("Greeter interface not implemeted")
	}
	return &greeterServer{Broker: b, IGreeter: p.F()}, nil
}

////////////////////////////////////////////////////////////////////////////////
//
//  For each function defined in the interface, there needs to be an equivalent
//  greeterServer call
//
////////////////////////////////////////////////////////////////////////////////

// Name calls the plugin implementation to get the name of the plugin
func (p *greeterServer) Name(nothing interface{}, result *string) error {
	*result = p.IGreeter.Name()
	return nil
}

// Version calls the plugin implementation to get the version of the plugin
func (p *greeterServer) Version(nothing interface{}, result *string) error {
	*result = p.IGreeter.Version()
	return nil
}

// StartTime calls the plugin implementation to get the start time of the plugin
func (p *greeterServer) StartTime(nothing interface{}, result *int64) error {
	*result = p.IGreeter.StartTime()
	return nil
}
