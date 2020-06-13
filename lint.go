package main

import (
	"fmt"
	"io/ioutil"
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

func lintTransactionFile(filePath, accountsPath string, outputJSON bool) {
	transactionsStr, err := readFileContent(filePath)
	if err != nil {
		panic(err)
	}

	transaction := Transaction{headerIdx: 1}
	validator := newValidator(filePath, accountsPath, outputJSON)
	iLine := 0

	for _, line := range strings.Split(transactionsStr, "\n") {
		iLine++

		// When the line is empty, skip it
		if commentOrEmptyPattern.MatchString(line) {
			continue
		}

		// When the line is a transaction header, validate and clear transaction
		transactionNext, headerParseError := parseTransactionHeader(iLine, line)
		if headerParseError == nil {
			validator.checkBalancingAndAccounts(transaction)
			transaction = transactionNext
			continue
		}

		// When the line is a posting, append it to transaction.postings
		postingParseSucceed, posting := parsePostingStr(line)
		if postingParseSucceed {
			transaction.postings = append(transaction.postings, posting)
			continue
		}

		if transaction.date == "" {
			// When the line is neither header or posting, return "Header unmatched" for compatibility
			validator.printer.warnHeaderUnmatched(transaction.headerIdx)
		} else {
			validator.printer.warnPostingParse(transaction.headerIdx, line)
		}
	}

	validator.checkBalancingAndAccounts(transaction)
}
