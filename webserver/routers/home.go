package routers
import (
	"net/http"
	"fmt"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/mcmx73/vaulta/webserver/forest"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Vaulta API")
}

func init() {
	log.Info("* Init router:home page")
	forest.AddRouterFunc("/i", Index)
}