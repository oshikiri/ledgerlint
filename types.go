package main

// See https://hledger.org/add.html#what-s-in-a-transaction

// Transaction contains its meta data and array of posting
type Transaction struct {
	date        string
	status      string
	description string
	postings    []Posting
}

// Posting contains its account type and amount object
type Posting struct {
	account  string
	figure   int
	currency string
}
