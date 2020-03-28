package main

// See https://hledger.org/add.html#what-s-in-a-transaction

// Date is date string
type Date string

// TransactionStatus should be one of ['*', '!', '']
type TransactionStatus string

// Amount for transaction posting, without currency
type Amount int

// Transaction contains its meta data and array of posting
type Transaction struct {
	date        Date
	status      TransactionStatus
	description string
	postings    []Posting
}

// calculateTotalAmount returns (containsOneEmptyAmount, totalAmount)
func (tx *Transaction) calculateTotalAmount() (bool, map[string]Amount) {
	totalAmounts := map[string]Amount{}
	containsOneEmptyAmount := false
	for _, posting := range tx.postings {
		if posting.emptyAmount {
			if containsOneEmptyAmount {
				// TODO: error contains two or more empty transactions
				return false, nil
			}
			containsOneEmptyAmount = true
		} else {
			totalAmounts[posting.currency] += posting.amount
		}
	}
	return containsOneEmptyAmount, totalAmounts
}

// Posting contains its account type and amount object
type Posting struct {
	account     string
	emptyAmount bool
	amount      Amount
	currency    string
}
