package stockItem

import (
	"fmillone/routes"
	"fmillone/utils"
	"net/http"

	"encoding/json"

	"fmt"

	"io"
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
	model := item.ToModel()

	if err := model.Save(); err != nil {
		panic(err)

	}

	m := utils.Message{"Model saved"}
	encodedMessage := utils.ToJSONOrPanic(m)
	if _, err := fmt.Fprint(w, encodedMessage); err != nil {
		panic(err)
	}
}

func GetRoutes() []routes.Route {
	return []routes.Route{
		routes.Route{
			utils.Post,
			"/item",
			createStockItem,
		},
	}
}
