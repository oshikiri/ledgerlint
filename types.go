package main

// See https://hledger.org/add.html#what-s-in-a-transaction

// Transaction contains its meta data and array of posting
type Transaction struct {
	date        string
	status      string
	description string
	postings    []Posting
}

func (tx *Transaction) calculateTotalAmount() map[string]int {
	totalAmounts := map[string]int{}
	for _, posting := range tx.postings {
		totalAmounts[posting.currency] += posting.figure
	}
	return totalAmounts
}

// Posting contains its account type and amount object
type Posting struct {
	account  string
	figure   int
	currency string
}
