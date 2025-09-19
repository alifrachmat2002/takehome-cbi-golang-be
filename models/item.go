package models

type Item struct {
	ID    int    `json:"id" binding:"gte=0"`
	Name  string `json:"name" binding:"required,min=1,max=100"`
	Price int    `json:"price" binding:"required,gt=0"`
}
