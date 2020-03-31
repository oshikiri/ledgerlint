package main

import (
	"fmt"
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
	// FIXME: 1000 USD + -1800 JPY
	msg := fmt.Sprintf(
		imbalancedTransactionMsg,
		filePath,
		lineNumber,
		strings.Join(amountStrs, " + "),
	)
	return msg
}
