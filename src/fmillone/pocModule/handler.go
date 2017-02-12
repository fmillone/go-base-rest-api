package pocModule

import (
	"fmillone/routes"
	"fmillone/utils"
	"fmt"
	"io"
	"net/http"

	"encoding/json"

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

func extractView(body io.Reader) *ModelDTO {
	item := ModelDTO{}
	err := json.NewDecoder(body).Decode(&item)
	if err != nil {
		fmt.Println("error", err)
	} else {
		return &item
	}
	return nil
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	modelDto := extractView(r.Body)
	model := Model{
		ID:   modelDto.ID,
		Name: modelDto.Name,
	}

	if err := model.Save(); err != nil {
		panic(err)

	}

	m := message{"Model saved"}
	encodedMessage := utils.ToJSONOrPanic(m)
	if _, err := fmt.Fprint(w, encodedMessage); err != nil {
		panic(err)
	}
}

func GetRoutes() []routes.Route {
	return []routes.Route{
		routes.Route{
			utils.Get,
			"/poc/:id",
			myHandler,
		},
		routes.Route{
			utils.Post,
			"/poc",
			saveHandler,
		},
	}
}
