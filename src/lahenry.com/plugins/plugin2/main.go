package main

import (
	"log"

	"time"

	"lahenry.com/library/baseplugin"
	"lahenry.com/library/config"
	"lahenry.com/plugins/plugin2/basepluginimpl"
)

func main() {
	flags := config.GetFlags()

	pluginOpts := &baseplugin.PluginOpts{
		Greeter:     basepluginimpl.Greeter,
		RunAsPlugin: flags.RunAsPlugin,
	}

	quit := make(chan bool)

	baseplugin.StartPlugin(pluginOpts, quit)

	go doSomething()

	<-quit // blocking
}

func doSomething() {
	greeter := basepluginimpl.Greeter()

	for idx := 0; true; idx++ {
		log.Printf("%s - %d", greeter.Name(), idx)

		time.Sleep(1000 * time.Millisecond)
	}
}
