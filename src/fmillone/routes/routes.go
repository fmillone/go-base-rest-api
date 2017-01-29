package routes

import (
	"net/http"

	"fmillone/utils"

	"github.com/go-zoo/bone"
)

var mux *bone.Mux

func SetMux(newMux *bone.Mux) {
	mux = newMux
}

//AddNewRoute adds a new route to the current mux
func AddNewRoute(method utils.HttpMethod, pattern string, handler http.HandlerFunc) {
	mux.Register(
		method.GetCode(),
		pattern,
		handler,
	)
}
