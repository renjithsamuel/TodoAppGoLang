package models

type Todo struct {
    ID        uint32   `json:"id"`
    Title     string   `json:"title"     validate:"required,min=3,max=50"`
    Completed bool     `json:"completed" validate:"required,boolean"`
}