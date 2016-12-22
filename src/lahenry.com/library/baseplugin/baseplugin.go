/*
Package baseplugin provides the bootstrap for Hashicorp go-plugin.

This package is used as the interface contract for the mainapp and plugins,
as well as plugin verification via handshake configuration.

Read these links for further information about Hashicorp go-plugin and its use:

https://github.com/hashicorp/go-plugin
https://godoc.org/github.com/hashicorp/go-plugin
https://www.terraform.io/docs/plugins/basics.html

[Excerpt from Usage section of https://github.com/hashicorp/go-plugin]

To use the plugin system, you must take the following steps. These are
high-level steps that must be done. Examples are available in the examples/
directory.

1. Choose the interface(s) you want to expose for plugins.

2. For each interface, implement an implementation of that interface that
communicates over an *rpc.Client (from the standard net/rpc package) for every
function call. Likewise, implement the RPC server struct this communicates to
which is then communicating to a real, concrete implementation.

3. Create a Plugin implementation that knows how to create the RPC client/server
for a given plugin type.

4. Plugin authors call plugin.Serve to serve a plugin from the main function.

5. Plugin users use plugin.Client to launch a subprocess and request an
interface implementation over RPC.

That's it! In practice, step 2 is the most tedious and time consuming step. Even
so, it isn't very difficult and you can see examples in the examples/ directory
as well as throughout our various open source projects.
*/

package baseplugin
