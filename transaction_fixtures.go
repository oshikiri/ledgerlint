package main

var transactionsBalanced = Transaction{
	date:        "2020-03-26",
	status:      "*",
	description: "super market",
	postings: []Posting{
		{
			account:  "Expences:Household essentials",
			amount:   200,
			currency: "JPY",
		},
		{
			account:  "Expences:Food",
			amount:   600,
			currency: "JPY",
		},
		{
			account:  "Assets:Cash",
			amount:   -800,
			currency: "JPY",
		},
	},
}

var transactionsBalancedEmptyAmount = Transaction{
	date:        "2020-03-26",
	status:      "*",
	description: "super market",
	postings: []Posting{
		{
			account:  "Expences:Household essentials",
			amount:   200,
			currency: "JPY",
		},
		{
			account:  "Expences:Food",
			amount:   600,
			currency: "JPY",
		},
		{
			account:     "Assets:Cash",
			emptyAmount: true,
		},
	},
}

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
			currency: "JPY",
		},
		{
			account:  "Expences:Household essentials",
			amount:   200,
			currency: "USD",
		},
		{
			account:  "Assets:Cash",
			amount:   -2000,
			currency: "USD",
		},
	},
}

var transactionDollar = Transaction{
	date: "2020/03/26",
	postings: []Posting{
		{
			account:  "Expences:Household essentials",
			amount:   1.23,
			currency: "$",
		},
		{
			account:  "Assets:Cash",
			amount:   -1.24,
			currency: "$",
		},
	},
}
