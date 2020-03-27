package main

var transactionsImbalanced = Transaction{
	date:        "2020-03-26",
	status:      "*",
	description: "toilet paper",
	postings: []Posting{
		{
			account:  "Expences:Household essentials",
			amount:   200,
			currency: "JPY",
		},
		{
			account:  "Assets:Cash",
			amount:   -2000,
			currency: "JPY",
		},
	},
}
