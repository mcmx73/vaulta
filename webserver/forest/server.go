package forest

import (
	"github.com/gorilla/mux"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"net/http"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

func init() {
	log.Info("* Init Forest server")
	router = mux.NewRouter().StrictSlash(true)
}

func StartServer() (err error) {

	http_port, err := config.GetString("HTTP_PORT")
	if err != nil {
		log.Fatal("Can not start HTTP server:", err)
		return
	}
	err = http.ListenAndServe(":" + http_port, router)
	return
}
