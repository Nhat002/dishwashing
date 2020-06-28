package main

import "errors"

const (
	// spongeMaxLiquid tells maximum number of liquid a sponge can contains
	spongeMaxLiquid uint = 10

	// Yellow is the constant for yellow color name
	Yellow string = "yellow"
	// Orange is the constant for orange color name
	Orange string = "orange"
)

// Sponge encapsulates attributes of the sponge
type Sponge struct {
	// Sponge has color, I prefer 3M yellow one
	Color string
	// washableTimes is the number of times we can use the Sponge to wash before toping up liquid
	washableTimes uint
}

// Wash washes the dish
func (s *Sponge) Wash(dish *Dish) error {
	if s.washableTimes == 0 {
		return errors.New("sponge has no more liquid for washing, topup please!")
	}
	s.washableTimes -= 1
	dish.clean()
	return nil
}

// Topup topups the liquid to the sponge
func (s *Sponge) Topup(liquid *LiquidDispenser) {
	topupAmount := spongeMaxLiquid
	if liquid.Amount() < topupAmount {
		topupAmount = liquid.Amount()
	}
	liquid.dispense(topupAmount)
	s.washableTimes += topupAmount
}
