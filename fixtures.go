package main

var transactionsImbalanced = Transaction{
	date:        "2020-03-26",
	status:      "*",
	description: "toilet paper",
	postings: []Posting{
		Posting{
			account:  "Expences:Household essentials",
			amount:   200,
			currency: "JPY",
		},
		Posting{
			account:  "Assets:Cash",
			amount:   -2000,
			currency: "JPY",
		},
	},
}
