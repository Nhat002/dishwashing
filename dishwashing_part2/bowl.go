package main

// Bowl encapsulates attributes of the Bowl
type Bowl struct {
	// dirty specifies the state of the Bowl, dirty or clean
	dirty bool
}

func NewBowl(dirty bool) *Bowl {
	return &Bowl{dirty: dirty}
}

// Clean makes the Bowl "not" dirty anymore
func (b *Bowl) clean() {
	b.dirty = false
}

// IsClean tells the state of the Bowl is clean or not
func (b Bowl) IsClean() bool {
	return !b.dirty
}
