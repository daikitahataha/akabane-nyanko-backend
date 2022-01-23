package models

type Cat struct {
	ID          uint   `json:"id"`
	name        string `json:"name"`
	description string `json:"description"`
}
