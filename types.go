package main

import "errors"

// See https://hledger.org/add.html#what-s-in-a-transaction

// Date is date string
type Date string

// TransactionStatus should be one of ['*', '!', null]
type TransactionStatus string

// Amount for transaction posting, without currency
type Amount float64

// Transaction contains its meta data and array of posting
type Transaction struct {
	date        Date
	status      TransactionStatus
	description string
	postings    []Posting
	headerIdx   int
}

// calculateTotalAmount returns (containsOneEmptyAmount, totalAmount)
func (tx *Transaction) calculateTotalAmount() (bool, map[string]Amount, int, error) {
	totalAmounts := map[string]Amount{}
	containsOneEmptyAmount := false
	emptyAmountLine := tx.headerIdx
	for _, posting := range tx.postings {
		if posting.emptyAmount {
			if containsOneEmptyAmount {
				return containsOneEmptyAmount, nil, emptyAmountLine, errors.New("Transaction contains two or more empty amount")
			}
			containsOneEmptyAmount = true
			emptyAmountLine = posting.line
		} else {
			totalAmounts[posting.currency] += posting.amount
		}
	}
	return containsOneEmptyAmount, totalAmounts, emptyAmountLine, nil
}

// Posting contains its account type and amount object
type Posting struct {
	account     string
	emptyAmount bool
	amount      Amount
	currency    string
	line        int
}
