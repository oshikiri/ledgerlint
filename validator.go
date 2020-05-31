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

		validator.knownAccounts = map[string]bool{}
		for _, account := range strings.Split(knownAccountsStr, "\n") {
			validator.knownAccounts[account] = true
		}
	}

	return &validator
}

// Validator checks transaction and posting, and print warning message if error is found.
type Validator struct {
	filePath      string
	accountsPath  string
	knownAccounts map[string]bool // values are not used
}

func (validator *Validator) checkUnknownAccount(countNewlines int, posting Posting) {
	if len(validator.knownAccounts) > 0 {
		_, exists := validator.knownAccounts[posting.account]
		if !exists {
			validator.warnParseFailed(
				countNewlines,
				fmt.Errorf("unknown account: %v", posting.account),
			)
		}
	}
}

func (validator *Validator) checkBalancing(countNewlines int, transaction Transaction) {
	containsOneEmptyAmount, totalAmount, err := transaction.calculateTotalAmount()

	if err != nil {
		validator.warnParseFailed(countNewlines, err)
	} else if !(isZeroAmount(totalAmount) || containsOneEmptyAmount) {
		validator.warnParseFailed(
			countNewlines,
			fmt.Errorf(buildImbalancedTransactionMsg(totalAmount)),
		)
	}
}

func (validator *Validator) warnParseFailed(countNewlines int, err error) {
	parseFailedMsg := "%v:%v %v\n"
	fmt.Printf(parseFailedMsg, validator.filePath, countNewlines, err)
}
