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
	account string
	amount  Amount
}

// Amount contains amount of money and its currency
type Amount struct {
	figure   int
	currency string
}
