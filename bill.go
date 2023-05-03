package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// create a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}
