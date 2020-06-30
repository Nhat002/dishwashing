package main

// Washable is the interface for kitchen items that can be washed
type Washable interface {
	clean()
	IsClean() bool
}
