package stockItem

import (
	"fmillone/routes"
	"fmillone/utils"
	"net/http"

	"encoding/json"

	"fmt"

	"io"

	"github.com/go-zoo/bone"
)

func extractView(body io.Reader) *JSONView {
	item := JSONView{}
	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		fmt.Println("error", err)
	} else {
		return &item
	}
	return nil
}

func createStockItem(w http.ResponseWriter, r *http.Request) {
	item := extractView(r.Body)
	fmt.Println("item", item)
	fmt.Fprintln(w, item)
}

//LoadModule ...
func LoadModule(mux *bone.Mux) {
	routes.SetMux(mux)
	routes.AddNewRoute(
		utils.Post,
		"/item",
		createStockItem,
	)
}
