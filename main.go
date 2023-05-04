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
		fmt.Println("You chose a")

	case "s":
		fmt.Println("You chose s")

	case "t":
		fmt.Println("You chose t")

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
