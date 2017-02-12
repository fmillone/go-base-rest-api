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

type Route struct {
	Method  utils.HttpMethod
	Pattern string
	Handler http.HandlerFunc
}

//AddNewRoute adds a new route to the current mux
func AddNewRoute(method utils.HttpMethod, pattern string, handler http.HandlerFunc) {
	mux.Register(
		method.GetCode(),
		pattern,
		handler,
	)
}

func AddRoute(route Route) {
	mux.Register(
		route.Method.GetCode(),
		route.Pattern,
		route.Handler,
	)
}
