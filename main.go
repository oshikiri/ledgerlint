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
	imbalancedTransactionMsg := "%v:%v imbalanced transaction, (total amount) = %v"
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
	fileContent := string(bytes)

	countNewlines := 1
	transactionStrs := strings.Split(fileContent, "\n\n")
	for _, transactionStr := range transactionStrs {
		_, transaction := parseTransactionStr(transactionStr)
		imbalancedTransactionMsg := buildImbalancedTransactionMsg(
			*filePath,
			countNewlines,
			transaction.calculateTotalAmount(),
		)
		fmt.Println(imbalancedTransactionMsg)
		countNewlines += strings.Count(transactionStr, "\n") + 2
	}
}
