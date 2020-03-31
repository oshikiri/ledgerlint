package main

import (
	"flag"
	"fmt"
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
	flag.Parse()

	knownAccountsStr, _ := readFileContent("fixtures/accounts.txt") // FIXME: error handling
	knownAccounts := strings.Split(knownAccountsStr, "\n")

	countNewlines := 1
	transactionsStr, _ := readFileContent(*filePath) // FIXME: error handling
	for _, transactionStr := range strings.Split(transactionsStr, "\n\n") {
		_, transaction := parseTransactionStr(transactionStr)

		// Check the transaction is balanced or not
		containsOneEmptyAmount, totalAmount := transaction.calculateTotalAmount()
		if !(isZeroAmount(totalAmount) || containsOneEmptyAmount) {
			imbalancedTransactionMsg := buildImbalancedTransactionMsg(
				*filePath,
				countNewlines,
				totalAmount,
			)
			fmt.Println(imbalancedTransactionMsg)
		}

		// Check unknown account
		for i, posting := range transaction.postings {
			if !contains(knownAccounts, posting.account) {
				unknownAccountMsg := "%v:%v unknown account: %v\n"
				fmt.Printf(unknownAccountMsg, *filePath, countNewlines+i+1, posting.account)
			}
		}

		countNewlines += strings.Count(transactionStr, "\n") + 2
	}
}
