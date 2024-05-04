package main

import "fmt"
import "os"

func main() {
	currBal := 2000
	fmt.Println("Welcome to go bank")
	banking(currBal)
}
func writeToFile(bal int) {
	balString := fmt.Sprint(bal)
	os.WriteFile("bal.txt", []byte(balString), 0644)
}
func banking(currBal int) {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Current Balance is: ", currBal)
	} else if choice == 2 {
		var deposit int
		fmt.Scan(&deposit)
		currBal += deposit
	} else if choice == 3 {
		var withdrawl int
		fmt.Scan(&withdrawl)
		currBal -= withdrawl
	} else {
		fmt.Println("Final balance written to file")
		writeToFile(currBal)
		return
	}
	banking(currBal)
}
