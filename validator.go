package main

import (
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
func newValidator(filePath, accountsPath string) *Validator {
	validator := Validator{
		filePath:     filePath,
		accountsPath: accountsPath,
	}

	if accountsPath != "" {
		knownAccountsStr, err := readFileContent(accountsPath)
		if err != nil {
			panic(err)
		}

		validator.knownAccounts = strings.Split(knownAccountsStr, "\n")
	}

	return &validator
}

// Validator checks transaction and posting, and print warning message if error is found.
type Validator struct {
	filePath      string
	accountsPath  string
	knownAccounts []string
}

func (validator *Validator) checkUnknownAccount(countNewlines int, posting Posting) {
	if len(validator.knownAccounts) > 0 && !contains(validator.knownAccounts, posting.account) {
		unknownAccountMsg := "%v:%v unknown account: %v\n"
		fmt.Printf(unknownAccountMsg, validator.filePath, countNewlines, posting.account)
	}
}

func (validator *Validator) checkBalancing(countNewlines int, transaction Transaction) {
	containsOneEmptyAmount, totalAmount, err := transaction.calculateTotalAmount()

	if err != nil {
		fmt.Printf("%v:%v %v\n", validator.filePath, countNewlines, err)
	} else if !(isZeroAmount(totalAmount) || containsOneEmptyAmount) {
		imbalancedTransactionMsg := buildImbalancedTransactionMsg(
			validator.filePath,
			countNewlines,
			totalAmount,
		)
		fmt.Println(imbalancedTransactionMsg)
	}
}

func (validator *Validator) warnParseFailed(countNewlines int, err error) {
	parseFailedMsg := "%v:%v %v\n"
	fmt.Printf(parseFailedMsg, validator.filePath, countNewlines, err)
}
