package main

// See https://hledger.org/add.html#what-s-in-a-transaction

type Date string
type TransactionStatus string
type Amount int

// Transaction contains its meta data and array of posting
type Transaction struct {
	date        Date
	status      TransactionStatus
	description string
	postings    []Posting
}

func (tx *Transaction) calculateTotalAmount() map[string]Amount {
	totalAmounts := map[string]Amount{}
	for _, posting := range tx.postings {
		totalAmounts[posting.currency] += posting.amount
	}
	return totalAmounts
}

// Posting contains its account type and amount object
type Posting struct {
	account  string
	amount   Amount
	currency string
}
