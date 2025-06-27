package server

type PayloadPostCreature struct {
	ID    string `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Race  string `json:"race" validate:"required"`
	Class string `json:"class" validate:"required"`
}
