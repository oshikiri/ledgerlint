package main

// See https://hledger.org/add.html#what-s-in-a-transaction

type Transaction struct {
	date        string
	status      string
	description string
	postings    []Posting
}

type Posting struct {
	account string
	amount  Amount
}

type Amount struct {
	figure   int
	currency string
}
