package main

func generateWashableList(number int) []Washable {
	washable := make([]Washable, number)
	for i := 0; i < number; i++ {
		washable[i] = generateWashableItem(i)
	}
	return washable
}

func generateWashableItem(number int) Washable {
	switch number % 4 {
	case 0:
		return NewDish(true)
	case 1:
		return NewBowl(true)
	case 2:
		return NewPan(true)
	case 3:
		return NewPot(true)
	default:
		return NewDish(true)
	}
}
