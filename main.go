package main

import (
	"flag"
	"strings"
)

func isZeroAmount(amounts map[string]Amount) bool {
	for _, v := range amounts {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	var filePath = flag.String("f", "", "ledger/hledger transaction file")
	var accountsPath = flag.String("account", "", "known accounts file")
	flag.Parse()

	validator := newValidator(*filePath, *accountsPath)

	countNewlines := 1
	transactionsStr, _ := readFileContent(*filePath) // FIXME: error handling
	for _, transactionStr := range strings.Split(transactionsStr, "\n\n") {
		_, transaction := parseTransactionStr(transactionStr)

		validator.checkBalancing(countNewlines, transaction)

		for i, posting := range transaction.postings {
			validator.checkUnknownAccount(countNewlines+i+1, posting)
		}

		countNewlines += strings.Count(transactionStr, "\n") + 2
	}
}
