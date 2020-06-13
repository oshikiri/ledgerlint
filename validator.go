package main

import (
	"fmt"
	"sort"
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
	validator := Validator{
		filePath:     filePath,
		accountsPath: accountsPath,
		outputJSON:   outputJSON,
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
	outputJSON    bool
	knownAccounts map[string]bool // values are not used
}

func (validator *Validator) checkUnknownAccount(countNewlines int, posting Posting) {
	if len(validator.knownAccounts) > 0 {
		_, exists := validator.knownAccounts[posting.account]
		if !exists {
			validator.warnUnknownAccount(countNewlines, posting.account)
		}
	}
}

func (validator *Validator) checkBalancing(countNewlines int, transaction Transaction) {
	containsOneEmptyAmount, totalAmount, err := transaction.calculateTotalAmount()

	if err != nil {
		validator.warnParseFailed(countNewlines, "ERROR", err)
	} else if !(isZeroAmount(totalAmount) || containsOneEmptyAmount) {
		validator.warnParseFailed(
			countNewlines,
			"ERROR",
			fmt.Errorf("imbalanced transaction, (total amount) = %v", calculateTotalAmount(totalAmount)),
		)
	}
}

func (validator *Validator) warnUnknownAccount(countNewlines int, account string) {
	validator.warnParseFailed(
		countNewlines,
		"WARN",
		fmt.Errorf("unknown account: %v", account),
	)
}

func (validator *Validator) warnHeaderUnmatched(countNewlines int) {
	validator.warnParseFailed(
		countNewlines,
		"WARN",
		fmt.Errorf("Header unmatched"),
	)
}

func (validator *Validator) warnPostingParse(countNewlines int, line string) {
	validator.warnParseFailed(
		countNewlines,
		"WARN",
		fmt.Errorf("parsePostingStr is failed: '%v'", line),
	)
}

func (validator *Validator) warnParseFailed(countNewlines int, logLevel string, err error) {
	if validator.outputJSON {
		parseFailedMsg := `{"file_path":"%v","line_number":%v,"level":"%v","error_message":"%v"}` + "\n"
		fmt.Printf(parseFailedMsg, validator.filePath, countNewlines, logLevel, err)
	} else {
		parseFailedMsg := "%v:%v %v\n"
		fmt.Printf(parseFailedMsg, validator.filePath, countNewlines, err)
	}
}
