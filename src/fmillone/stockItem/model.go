package stockItem

//StockItem base data type
type StockItem struct {
	ID          int64
	Name        string
	Quantity    uint64
	Description string
}

//JSONView : DTO for marshaling
type JSONView struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name"`
	Quantity    uint64 `json:"quantity"`
	Description string `json:"description"`
}

//NewJSONView constructs a JSONView
func NewJSONView(item StockItem) JSONView {
	return JSONView{
		ID:          item.ID,
		Name:        item.Name,
		Quantity:    item.Quantity,
		Description: item.Description,
	}
}

func (model StockItem) Save() error {
	return GetDatastore().Save(&model)
}

func (item JSONView) ToModel() StockItem {
	return StockItem{
		ID:          item.ID,
		Name:        item.Name,
		Quantity:    item.Quantity,
		Description: item.Description,
	}
}
