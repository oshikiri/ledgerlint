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

func (tx *Transaction) calculateTotalAmount() map[string]Amount {
	totalAmounts := map[string]Amount{}
	containsEmptyAmount := false
	for _, posting := range tx.postings {
		if posting.emptyAmount {
			if containsEmptyAmount {
				// TODO: error contains two or more empty transactions
				return nil
			} else {
				containsEmptyAmount = true
			}
		} else {
			totalAmounts[posting.currency] += posting.amount
		}
	}
	return totalAmounts
}

// Posting contains its account type and amount object
type Posting struct {
	account     string
	emptyAmount bool
	amount      Amount
	currency    string
}
