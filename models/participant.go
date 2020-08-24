package models

const (
	Small      = 'S'
	Medium     = 'M'
	Large      = 'L'
	ExtraLarge = "XL"
)

type Participant struct {
	Id         int
	Name       string
	Email      string
	Photo      string
	Phone      string
	TShirtSize string // TODO enum?
}
