package main

import (
	"fmt"
	"reflect"
)

func main() {
	washableItems := generateWashableList(10)
	sponge := &Sponge{Color: Yellow}
	liquidDispenser := NewLiquidDispenser(100)

	sponge.Topup(liquidDispenser)
	for i, item := range washableItems {
		itemType := reflect.TypeOf(item).Elem().Name()
		fmt.Printf("item at index %d is type %s\n", i+1, itemType)
		fmt.Printf("Before wash, is %s clean? = %t\n", itemType, item.IsClean())
		sponge.Wash(item)
		fmt.Printf("After wash, is %s clean? = %t\n", itemType, item.IsClean())
	}
}
