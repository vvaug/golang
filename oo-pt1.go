package main

import (
	"fmt"
	"os"
)

type Costumer struct {
	name     string
	document string
}

type CurrentAccount struct {
	issuer        int
	accountNumber int
	balance       float64
	costumer      *Costumer
}

var costumers []Costumer
var accounts []CurrentAccount

func main() {
	for {
		option := option()
		proccess(option)
	}
}

func option() int {
	fmt.Println("1) - Create an account")
	fmt.Println("2) - Account operations")
	fmt.Println("3) - Exit")
	var option int
	fmt.Scan(&option)
	return option
}

func proccess(option int) {
	switch option {
	case 1:
		fmt.Println("document number:")
		var documentNumber string
		fmt.Scan(&documentNumber)
		if isValidCustomer(documentNumber) {
			costumer := getCostumer(documentNumber)
			var issuer int
			var accountNumber int
			fmt.Println("issuer:")
			fmt.Scan(&issuer)
			fmt.Println("number:")
			fmt.Scan(&accountNumber)
			account := CurrentAccount{issuer, accountNumber, 0, costumer}
			accounts = append(accounts, account)
		} else {
			fmt.Println("You are not costumer yet. Please, informe your personal data")
			createCostumer()
		}
	case 2:
		fmt.Println("document number:")
		var documentNumber string
		fmt.Scan(&documentNumber)
		fmt.Println("account number:")
		var accountNumber int
		fmt.Scan(&accountNumber)
		if isValidCostumerAndAccount(documentNumber, accountNumber) {
			account := getAccount(accountNumber)
			resume(account)
			operations(account)
		} else {
			panic("Costumer or account data is invalid.")
		}
	case 3:
		os.Exit(0)
	}
}

func operations(account *CurrentAccount) {
	fmt.Println("1) - Withdrawal")
	fmt.Println("2) - Deposit")
	fmt.Println("3) - History")
	fmt.Println("4) - Exit")
	var operation int
	fmt.Scan(&operation)
	switch operation {
	case 1:
		fmt.Println("amount:")
		var amount float64
		fmt.Scan(&amount)
		withdrawal(account, amount)

	case 2:
		fmt.Println("amount:")
		var amount float64
		fmt.Scan(&amount)
		deposit(account, amount)

	case 3:
		fmt.Println("Not implemented")

	case 4:
		os.Exit(0)
	}
}

func deposit(account *CurrentAccount, amount float64) {
	account.balance = account.balance + amount
	fmt.Println("Your new balance:", account.balance)
}

func withdrawal(account *CurrentAccount, amount float64) {
	if account.balance == 0 || account.balance < amount {
		panic("Account has no balance")
	}
	account.balance = account.balance - amount
	fmt.Println("Your new balance:", account.balance)
}

func resume(account *CurrentAccount) {
	fmt.Println("Your data was been validated.")
	fmt.Println()
	fmt.Println("Account issuer:", account.issuer)
	fmt.Println("Account number:", account.accountNumber)
	fmt.Println("Account balance:", account.balance)
	fmt.Println("Costumer:", account.costumer.name)
}

func isValidCostumerAndAccount(documentNumber string, accountNumber int) bool {
	account := getAccount(accountNumber)
	return account.costumer.document == documentNumber
}

func getAccount(accountNumber int) *CurrentAccount {
	for i := 0; i < len(accounts); i++ {
		if accounts[i].accountNumber == accountNumber {
			return &accounts[i]
		}
	}
	panic("Doesn't exist any account with the entered number")
}

func getCostumer(document string) *Costumer {
	for i := 0; i < len(costumers); i++ {
		if costumers[i].document == document {
			return &costumers[i]
		}
	}
	panic("Costumer doesn't exist.")
}

func isValidCustomer(document string) bool {

	if len(costumers) == 0 {
		return false
	}

	for i := 0; i < len(costumers); i++ {
		if costumers[i].document == document {
			return true
		}
	}
	return false
}

func createCostumer() Costumer {
	var name string
	var document string
	fmt.Println("name:")
	fmt.Scan(&name)
	fmt.Println("document:")
	fmt.Scan(&document)
	costumer := Costumer{name, document}
	costumers = append(costumers, costumer)
	fmt.Println("Costumer:", costumer.document, "registered successfully.")
	return costumer
}
