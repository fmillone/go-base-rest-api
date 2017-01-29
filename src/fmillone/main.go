package main

import (
	"fmillone/pocModule"

	"github.com/codegangsta/negroni"
	"github.com/go-zoo/bone"
)

func buildMux() *bone.Mux {
	mux := bone.New()
	pocModule.LoadModule(mux)

	return mux
}

func main() {
	n := negroni.Classic()

	n.UseHandler(buildMux())
	n.Run(":3000")
}
