package main

// Pan encapsulates attributes of the Pan
type Pan struct {
	// dirty specifies the state of the Pan, dirty or clean
	dirty bool
}

func NewPan(dirty bool) *Pan {
	return &Pan{dirty: dirty}
}

// Clean makes the Pan "not" dirty anymore
func (b *Pan) clean() {
	b.dirty = false
}

// IsClean tells the state of the Pan is clean or not
func (b Pan) IsClean() bool {
	return !b.dirty
}
