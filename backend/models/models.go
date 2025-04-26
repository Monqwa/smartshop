package models

import "github.com/google/uuid"

type ItemModel struct{}
type Item struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	IsBuy  bool      `json:"is_buy"`
	ListID uuid.UUID `json:"list_id"`
}

func (i *Item) TableName() string {
	return "items"
}

type ListModel struct{}
type List struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}

func (l *List) TableName() string {
	return "list"
}
