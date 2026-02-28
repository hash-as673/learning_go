package main

import "fmt"

func processTransaction(amount int) (status string) {
	defer func() {
		fmt.Println("Audit Log: Transaction processed with final status:", status)
	}()
	if amount < 0 {
		status = "Rejected: Negative Amount"
		return
	} else if amount > 10000 {
		status = "Flagged for Review"
		return
	} else {
		status = "Approved"
		return
	}
}

func main() {
	processTransaction(9990)
	processTransaction(11000)
	processTransaction(-20)
}
