package routers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/mcmx73/vaulta/webserver/forest"
	"github.com/flosch/pongo2"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var ping_answer map[string]interface{}
	ping_answer = make(map[string]interface{})
	ping_answer["ping"] = "pong"
	ping_answer["service"] = "Vaulta API"
	ping_answer["result_code"] = 200
	ping_string, _ := json.Marshal(ping_answer)
	fmt.Fprintf(w, string(ping_string))
}

func init() {
	log.Info("* Init base API")
	tplExample = pongo2.Must(pongo2.FromFile("view/default.tpl"))
	forest.AddRouterFunc("/api/ping", Ping)
	forest.AddRouterFunc("/", ExamplePage)
}

var tplExample *pongo2.Template

func ExamplePage(w http.ResponseWriter, r *http.Request) {
	// Execute the template per HTTP request
	err := tplExample.ExecuteWriter(pongo2.Context{"query": r.FormValue("query")}, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}