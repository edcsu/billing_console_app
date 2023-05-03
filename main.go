package main

import "fmt"

func main() {
	todayBill := newBill("John's bill")

	todayBill.addItem("coffee", 5000.25)
	todayBill.addItem("samosa", 500.50)
	todayBill.addItem("cake slice", 2500.75)

	fmt.Println(todayBill.format())
}
