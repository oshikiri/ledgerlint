package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func buildImbalancedTransactionMsg(
	filePath string,
	lineNumber int,
	amounts map[string]Amount,
) string {
	amountStrs := []string{}
	for currency, amount := range amounts {
		amountAndCurrency := fmt.Sprintf("%v %v", amount, currency)
		amountStrs = append(amountStrs, amountAndCurrency)
	}
	imbalancedTransactionMsg := "%v:%v imbalanced transaction, (total amount) = %v"
	msg := fmt.Sprintf(
		imbalancedTransactionMsg,
		filePath,
		lineNumber,
		strings.Join(amountStrs, " + "),
	)
	return msg
}

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

	bytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}
	fileContent := string(bytes)

	countNewlines := 1
	transactionStrs := strings.Split(fileContent, "\n\n")
	for _, transactionStr := range transactionStrs {
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

		countNewlines += strings.Count(transactionStr, "\n") + 2
	}
}
