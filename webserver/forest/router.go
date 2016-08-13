package forest

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/DeepForestTeam/mobiussign/components/log"
)

var router = mux.NewRouter().StrictSlash(false)

func init() {
	log.Info("* Init Forest router")
}

func AddRouterFunc(uri string, f func(http.ResponseWriter, *http.Request)) {
	log.Debug("Register flat handler for:", uri)
	router.HandleFunc(uri, f)
}
func AddRouter(uri string, base_handler Controller) {
	router.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		handler := base_handler
		handler.Process(w, r)
		stop_process := handler.PreRoute()
		if !stop_process {
			switch r.Method {
			case "GET":
				handler.Get()
			case "POST":
				handler.Post()
			case "PUT":
				handler.Put()
			case "DELETE":
				handler.Delete()
			case "OPTIONS":
				handler.Options()
			}
		}
		handler.RenderTemplate()
	})
}