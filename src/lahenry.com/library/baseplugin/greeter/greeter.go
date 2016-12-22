// Package greeter defines the interface contract between Hashicorp go-plugin
// for the RPC sender and receiver for the greeter interface.

package greeter

// InterfaceName is the unique name for this interface
const InterfaceName = "Greeter"

// IGreeter is the interface that is known between the mainapp and plugins
type IGreeter interface {
	Name() string
	Version() string
	StartTime() int64
}

// Greeter is a structure that has a reference to the interface
type Greeter struct {
	F func() IGreeter
}

// GreeterFunc is the callback function
type GreeterFunc func() IGreeter
