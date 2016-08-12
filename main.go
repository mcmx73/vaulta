package main

import (
	"github.com/DeepForestTeam/mobiussign/components/log"
	_ "github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/mcmx73/vaulta/components/store"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
)

func init() {
	log.Info("* Init main")
}
func main() {
	err := store.ConnectDB()
	if err != nil {
		log.Fatal("Can not connect to storage:", err)
		panic(err)
	}
	log.Debug("Bolt/StromDB connected")
	err = forest.StartServer()
	log.Fatal(err)
}
