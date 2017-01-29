package pocModule

import (
	"fmillone/routes"
	"fmillone/utils"
	"fmt"
	"net/http"

	"github.com/go-zoo/bone"
)

type message struct {
	Message string
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	m := message{
		fmt.Sprintf("Hello %v", bone.GetValue(r, "id")),
	}
	encodedMessage := utils.ToJSONOrPanic(m)
	if _, err := fmt.Fprint(w, encodedMessage); err != nil {
		panic(err)
	}
}

//LoadModule ...
func LoadModule(mux *bone.Mux) {
	routes.SetMux(mux)
	routes.AddNewRoute(
		utils.Get,
		"/poc/:id",
		myHandler,
	)
}
