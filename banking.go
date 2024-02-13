package main

import (
	"banking_app/filesys"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

// WRITING AND READING BALANCE TO AND FROM FILE, INSTEAD OF LOSING WHEN PROGRAM IS RESTARTED

const accountBalanceFile = "balance.txt"

// MAIN BANK LOGIC
func main() {
	var accountBalance, err = filesys.ReadFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("=================")
		//return
		//panic("Can't continue, sorry")
	}

	fmt.Println("Welcome to Go Bank of Nigeria")
	fmt.Println("You can reach us here 24/7:", randomdata.PhoneNumber())

	for {

		presentOptions()

		var choice int
		fmt.Print("What would you like to do today?: ")
		fmt.Scan(&choice)

		//wantsToCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your Balance is", accountBalance)
		case 2:
			var depositAmount float64
			fmt.Print("How much would you like to deposit: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid Amount. Amount must be greater than 0")
				//return
				// using the CONTINUE keyword below cuz I want code to keep
				//running instead of breaking
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Thank You!!")
			fmt.Println("Your new balance is", accountBalance)
			filesys.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			var withdrawAmount float64
			fmt.Print("How much would you like to withdraw: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 || withdrawAmount > accountBalance {
				fmt.Println("Invalid Amount. Amount must be greater than 0 and less than account balance.")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Thank You!!")
			fmt.Println("Your new balance is", accountBalance)
			filesys.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!!!")
			fmt.Println("Thanks for choosing Go Bank")
			return
		}

		// ABOVE I USED THE SWITCH STATEMENT INSTEAD OF IF - ELSE

		// if choice == 1 {
		// 	fmt.Println("Your Balance is", accountBalance)
		// } else if choice == 2 {
		// 	var depositAmount float64
		// 	fmt.Print("How much would you like to deposit: ")
		// 	fmt.Scan(&depositAmount)
		// 	if depositAmount <= 0 {
		// 		fmt.Println("Invalid Amount. Amount must be greater than 0")
		// 		continue
		// 	}
		// 	accountBalance += depositAmount
		// 	fmt.Println("Thank You!!")
		// 	fmt.Println("Your new balance is", accountBalance)
		// } else if choice == 3 {
		// 	var withdrawAmount float64
		// 	fmt.Print("How much would you like to withdraw: ")
		// 	fmt.Scan(&withdrawAmount)
		// 	if withdrawAmount <= 0 || withdrawAmount > accountBalance {
		// 		fmt.Println("Invalid Amount. Amount must be greater than 0 and less than account balance.")
		// 		continue
		// 	}
		// 	accountBalance -= withdrawAmount
		// 	fmt.Println("Thank You!!")
		// 	fmt.Println("Your new balance is", accountBalance)
		// } else {
		// 	fmt.Println("Goodbye!!!")
		//return
		//I chose to use the BREAK keyword in next line instead of RETURN
		//because I need to run an ending message outside the loop
		// 	break
		// }
	}
}
