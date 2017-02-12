package main

import (
	"fmillone/routes"
	"fmillone/stockItem"

	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
)

func buildMux() *bone.Mux {
	mux := bone.New()
	routes.SetMux(mux)
	for _, route := range stockItem.GetRoutes() {
		routes.AddRoute(route)
	}

	return mux
}

func main() {
	n := negroni.Classic()

	n.UseHandler(buildMux())
	n.Run(":3000")
}
