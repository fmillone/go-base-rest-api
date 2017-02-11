package main

import (
	"fmillone/pocModule"
	"fmillone/routes"

	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
)

func buildMux() *bone.Mux {
	mux := bone.New()
	routes.SetMux(mux)
	for _, route := range pocModule.GetRoutes() {
		routes.AddRoute(route)
	}

	return mux
}

func main() {
	n := negroni.Classic()

	n.UseHandler(buildMux())
	n.Run(":3000")
}
