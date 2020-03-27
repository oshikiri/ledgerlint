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

var transactionsImbalancedMultiCurrency = Transaction{
	date:        "2020-03-01",
	status:      "!",
	description: "at super market",
	postings: []Posting{
		{
			account:  "Expences:Food",
			amount:   1000,
			currency: "USD",
		},
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
