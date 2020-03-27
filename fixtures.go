package main

var transactionsImbalanced = Transaction{
	date:        "2020-03-26",
	status:      "*",
	description: "toilet paper",
	postings: []Posting{
		Posting{
			account:  "Expences:Household essentials",
			figure:   200,
			currency: "JPY",
		},
		Posting{
			account:  "Assets:Cash",
			figure:   -2000,
			currency: "JPY",
		},
	},
}
