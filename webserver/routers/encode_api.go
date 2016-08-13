package routers

import (
	"github.com/mcmx73/vaulta/webserver/forest"
	"github.com/mcmx73/vaulta/webserver/controls"
)

func init() {
	sign_api := controls.EncoderController{}
	sign_api.ThisName = "EncoderAPI"
	forest.AddRouter("/api", &sign_api)
	forest.AddRouter("/api/{data_link:[0-9A-Fa-f]{16}}/{data_key:[0-9A-Fa-f]{16}}", &sign_api)
}
