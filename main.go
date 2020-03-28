package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func buildImbalancedTransactionMsg(filePath string, lineNumber int, amounts map[string]Amount) string {
	amountStrs := []string{}
	for currency, amount := range amounts {
		amountAndCurrency := fmt.Sprintf("%v %v", amount, currency)
		amountStrs = append(amountStrs, amountAndCurrency)
	}
	imbalancedTransactionMsg := "%v:%v imbalanced transaction is found. Total amount = (%v)"
	msg := fmt.Sprintf(
		imbalancedTransactionMsg,
		filePath,
		lineNumber,
		strings.Join(amountStrs, " + "),
	)
	return msg
}

func main() {
	var filePath = flag.String("f", "", "ledger/hledger transaction file")
	flag.Parse()

	bytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}

	transactionStrs := strings.Split(string(bytes), "\n\n")
	_, transaction := parseTransactionStr(transactionStrs[0])
	lineNumber := 1
	imbalancedTransactionMsg := buildImbalancedTransactionMsg(
		*filePath,
		lineNumber,
		transaction.calculateTotalAmount(),
	)
	fmt.Println(imbalancedTransactionMsg)
}
