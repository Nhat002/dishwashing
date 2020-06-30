package main

// Dish encapsulates attributes of the dish
type Dish struct {
	// dirty specifies the state of the dish, dirty or clean
	dirty bool
}

func NewDish(dirty bool) *Dish {
	return &Dish{dirty: dirty}
}

// Clean makes the dish "not" dirty anymore
func (d *Dish) clean() {
	d.dirty = false
}

// IsClean tells the state of the dish is clean or not
func (d Dish) IsClean() bool {
	return !d.dirty
}
