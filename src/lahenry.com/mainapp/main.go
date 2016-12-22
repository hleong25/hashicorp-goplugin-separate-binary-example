package main

import (
	"log"
	"time"

	"lahenry.com/mainapp/pluginmgr"
)

func main() {

	var pluginMgr = &pluginmgr.PluginManager{}

	pluginMgr.StartPlugin("plugin1")
	pluginMgr.StartPlugin("plugin2")

	quit := make(chan bool)

	go doSomething()

	<-quit // blocking
}

func doSomething() {

	for idx := 0; true; idx++ {
		log.Printf("mainapp - %d", idx)

		time.Sleep(1000 * time.Millisecond)
	}
}
