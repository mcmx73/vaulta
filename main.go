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
	if tapeblock.TotalCoutunt() < 10000 {
		log.Debug("Generate random blocks")
		tapeblock.FillRandom(10000)
		log.Debug("Now total blocks:", tapeblock.TotalCoutunt())
	}
	tapeblock.BlockData = []byte("The Ultimate Question of Life, the Universe, and Everything:42....")
	offset, key, err := tapeblock.Encrypt()
	//8bc9d18124f1a18d 1190ec32be94eeb0
	decrypt_test := tapeencript.TapeBlock{BlockId:offset}
	decrypt_test.Decrypt(key)
	log.Warning("decrypt: [", string(decrypt_test.BlockData), "]")
	log.Debug(offset, key, err)
	err = forest.StartServer()
	log.Fatal(err)
}
