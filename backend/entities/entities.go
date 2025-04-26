package entities

type CreateItem struct {
	Name   string `json:"name" required:"true"`
	IsBuy  bool   `json:"is_buy" default:"false"`
	ListID string `json:"list_id" required:"true"`
}

type UpdateItem struct {
	Name  string `json:"name" required:"true"`
	IsBuy *bool  `json:"is_buy" default:"false"`
}

type CreateList struct {
	Title string `json:"title" required:"true"`
}

type UpdateList struct {
	Title string `json:"title" required:"true"`
}
