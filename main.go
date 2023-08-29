package main

import (
	"github.com/Prouk/moujinReloaded/src/core"
	"log"
)

func main() {

	m := new(core.Moujin)
	m.Config = new(core.Config)
	err := m.Config.SetDefaultConf()
	if err != nil {
		log.Fatalf("Error parsing config file : %s", err)
		return
	}
	err = m.SetDefaultViews()
	if err != nil {
		log.Fatalf("Error parsing html files : %s", err)
		return
	}
	m.SetDefaultRouter()
	m.Router.Run(":" + m.Config.Port)
}
