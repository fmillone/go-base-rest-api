package pocModule

type Model struct {
	ID   int
	Name string
}

type ModelDTO struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (model Model) Save() error {
	return GetDatastore().Save(&model)
}
