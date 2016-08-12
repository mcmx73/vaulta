package main

import (
	"github.com/DeepForestTeam/mobiussign/components/log"
	_ "github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/mcmx73/vaulta/components/store"
	"github.com/DeepForestTeam/mobiussign/restapi/forest"
	"github.com/mcmx73/vaulta/components/tapeencript"
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
	if tapeblock.TotalCoutunt()<10000{
		log.Debug("Generate random blocks")
		tapeblock.FillRandom(10000)
		log.Debug("Now total blocks:", tapeblock.TotalCoutunt())
	}
	err = forest.StartServer()
	log.Fatal(err)
}
