package models

type Menu struct {
	Breakfast      []Dish `json:"breakfast"`
	Lunch          []Dish `json:"lunch"`
	AfternoonSnack []Dish `json:"afternoon_snack"`
	Dinner         []Dish `json:"dinner"`
}

type Dish struct {
	Name        string `json:"name"`
	Ingredients string `json:"ingredients"`
	Recipe      string `json:"recipe"`
}
