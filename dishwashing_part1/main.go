package main

import (
	"fmt"
)

func main() {
	dishes := generateDirtyDishes(10)
	sponge := &Sponge{Color: Yellow}
	liquidDispenser := NewLiquidDispenser(100)

	sponge.Topup(liquidDispenser)
	for i, dish := range dishes {
		fmt.Printf("Before wash, is dish %d clean? = %t\n", i+1, dish.IsClean())
		sponge.Wash(dish)
		fmt.Printf("After wash, is dish %d clean? = %t\n", i+1, dish.IsClean())
	}
}

func generateDirtyDishes(amount uint) []*Dish {
	dishes := make([]*Dish, amount)
	for i := range dishes {
		dishes[i] = NewDish(true)
	}
	return dishes
}
