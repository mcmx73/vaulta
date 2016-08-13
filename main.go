package main

import (
	_ "github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/mcmx73/vaulta/components/store"
	"github.com/mcmx73/vaulta/webserver/forest"
	"github.com/mcmx73/vaulta/components/tapeencript"
	_ "github.com/mcmx73/vaulta/webserver/routers"
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
	tapeblock := tapeencript.TapeBlock{}
	log.Debug("Total blocks:", tapeblock.TotalCoutunt())
	//Initial random fill
	if tapeblock.TotalCoutunt() < 10000 {
		log.Debug("Generate random blocks")
		tapeblock.FillRandom(10000)
		log.Debug("Now total blocks:", tapeblock.TotalCoutunt())
	}
	err = forest.StartServer()
	log.Fatal(err)
}
