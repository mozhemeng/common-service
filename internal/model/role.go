package model

type Role struct {
	*BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
