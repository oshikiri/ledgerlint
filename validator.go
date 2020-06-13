package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func readFileContent(filePath string) (string, error) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ioutil.ReadFile failed: %v, filePath='%v'\n", err, filePath)
		return "", err
	}
	fileContent := string(bytes)
	return fileContent, nil
}

func isZeroAmount(amounts map[string]Amount) bool {
	for _, v := range amounts {
		if v != 0 {
			return false
		}
	}
	return true
}

func calculateTotalAmount(
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

	return strings.ReplaceAll(strings.Join(amountStrs, " + "), "+ -", "- ")
}

func newValidator(filePath, accountsPath string, outputJSON bool) *Validator {
	printer := Printer{
		filePath:   filePath,
		outputJSON: outputJSON,
	}
	validator := Validator{
		accountsPath: accountsPath,
		printer:      printer,
	}

	if accountsPath != "" {
		knownAccountsStr, err := readFileContent(accountsPath)
		if err != nil {
			panic(err)
		}

		validator.knownAccounts = map[string]bool{}
		for _, account := range strings.Split(knownAccountsStr, "\n") {
			validator.knownAccounts[account] = true
		}
	}

	return &validator
}

// Validator checks transaction and posting, and print warning message if error is found.
type Validator struct {
	accountsPath  string
	printer       Printer
	knownAccounts map[string]bool // values are not used
}

func (validator *Validator) checkUnknownAccount(countNewlines int, posting Posting) {
	if len(validator.knownAccounts) > 0 {
		_, exists := validator.knownAccounts[posting.account]
		if !exists {
			validator.printer.warnUnknownAccount(countNewlines, posting.account)
		}
	}
}

func (validator *Validator) checkBalancing(countNewlines int, transaction Transaction) {
	containsOneEmptyAmount, totalAmount, err := transaction.calculateTotalAmount()

	if err != nil {
		validator.printer.print(countNewlines, "ERROR", err)
	} else if !(isZeroAmount(totalAmount) || containsOneEmptyAmount) {
		validator.printer.print(
			countNewlines,
			"ERROR",
			fmt.Errorf("imbalanced transaction, (total amount) = %v", calculateTotalAmount(totalAmount)),
		)
	}
}

func (validator *Validator) checkBalancingAndAccounts(transaction Transaction) {
	transactionHeaderIdx := transaction.headerIdx
	validator.checkBalancing(transactionHeaderIdx, transaction)

	for i, posting := range transaction.postings {
		validator.checkUnknownAccount(transactionHeaderIdx+i+1, posting)
	}
}
