package main

import "fmt"

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue:")
	fmt.Scan(&revenue)

	fmt.Print("expenses:")
	fmt.Scan(&expenses)

	fmt.Print("taxRate:")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)

	fmt.Println("Earnings Before Tax:", ebt)
	fmt.Println("Net Profit:", profit)

}
