package main

import (
	"github.com/raghukul01/Remotify/config"
	"github.com/raghukul01/Remotify/init"
)

func main() {
	config.Load()
	webServer := server.New()
	webServer.ServeHTTP()
}
