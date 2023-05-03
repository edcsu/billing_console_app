package main

import "fmt"

func main() {
	todayBill := newBill("John's bill")

	fmt.Println(todayBill.format())
}
