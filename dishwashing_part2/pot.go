package main

// Pot encapsulates attributes of the Pot
type Pot struct {
	// dirty specifies the state of the Pot, dirty or clean
	dirty bool
}

func NewPot(dirty bool) *Pot {
	return &Pot{dirty: dirty}
}

// Clean makes the Pot "not" dirty anymore
func (b *Pot) clean() {
	b.dirty = false
}

// IsClean tells the state of the Pot is clean or not
func (b Pot) IsClean() bool {
	return !b.dirty
}
