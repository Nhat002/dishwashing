package main

import "errors"

const (
	liquidDispenserMaxAmount = 100
)

// LiquidDispenser encapsulates the attributes of dish washing liquid dispenser
type LiquidDispenser struct {
	// amount is the amount of liquid in the dish washing liquid bottle
	// to make it easier for the model, it is an usigned integer. And one liquid amount can be used to wash 1 dish.
	amount uint
}

func NewLiquidDispenser(amount uint) *LiquidDispenser {
	return &LiquidDispenser{amount}
}

// Amount shows the amount of liquid to public
func (liquid LiquidDispenser) Amount() uint {
	return liquid.amount
}

// Dispense an amount of liquid
func (liquid *LiquidDispenser) dispense(amount uint) {
	if amount > liquid.amount {
		amount = liquid.amount
	}
	liquid.amount = liquid.amount - amount
}

func (liquid *LiquidDispenser) Topup(amount uint) error {
	if amount+liquid.amount > liquidDispenserMaxAmount {
		return errors.New("cannot topup more than liquid dispenser's max amount")
	}

	liquid.amount += amount
	return nil
}
