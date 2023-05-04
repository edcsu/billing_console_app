package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose an option: \n a - add item \n s - save bill \n t - add tip: \n", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		fmt.Println(name, price)

	case "s":
		fmt.Println("You chose s")

	case "t":
		tip, _ := getInput("Enter tip amount (UGX): ", reader)

		fmt.Println(tip)

	default:
		fmt.Println("Invalid option...")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Type in the new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	todaysBill := createBill()
	promptOptions(todaysBill)
}
