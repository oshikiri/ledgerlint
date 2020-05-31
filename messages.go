package main

import (
	"fmt"
	"sort"
	"strings"
)

func buildImbalancedTransactionMsg(
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
	msg := fmt.Sprintf(
		"imbalanced transaction, (total amount) = %v",
		strings.ReplaceAll(strings.Join(amountStrs, " + "), "+ -", "- "),
	)
	return msg
}
