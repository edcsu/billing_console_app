package main

import (
	"fmt"
	"os"
)

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

// receiver function
func (b bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...... UGX%v \n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-25v ...... UGX%v \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...... UGX%0.2f", "Total:", total+b.tip)

	return fs
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) addTip(tip float64) {
	b.tip = tip
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file successfully")
}
