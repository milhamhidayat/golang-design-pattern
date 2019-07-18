package meal

type Meal struct {
	Chicken string
	Drink   string
	Option  string
}

type MealProduct interface {
	MakeMeal() MealProduct
}
