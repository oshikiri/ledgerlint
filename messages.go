package main

import (
	"fmt"
	"sort"
	"strings"
)

func buildImbalancedTransactionMsg(
	filePath string,
	lineNumber int,
	amounts map[string]Amount,
) string {
	currencies := make([]string, 0, len(amounts))
	for currency := range amounts {
		currencies = append(currencies, currency)
	}
	sort.Strings(currencies)

	amountStrs := []string{}
	for _, currency := range currencies {
		amount := amounts[currency]
		amountAndCurrency := fmt.Sprintf("%v %v", amount, currency)
		amountStrs = append(amountStrs, amountAndCurrency)
	}
	imbalancedTransactionMsg := "%v:%v imbalanced transaction, (total amount) = %v"
	msg := fmt.Sprintf(
		imbalancedTransactionMsg,
		filePath,
		lineNumber,
		strings.ReplaceAll(strings.Join(amountStrs, " + "), "+ -", "- "),
	)
	return msg
}
