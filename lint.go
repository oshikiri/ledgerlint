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

	var transaction Transaction
	validator := newValidator(filePath, accountsPath, outputJSON)
	transactionHeaderIdx := 1

	for iLine, line := range strings.Split(transactionsStr, "\n") {
		// When the line is empty, skip it
		if commentOrEmptyPattern.MatchString(line) {
			continue
		}

		// When the line is a transaction header, validate and clear transaction
		transactionNext, headerParseError := parseTransactionHeader(line)
		if headerParseError == nil {
			validator.checkBalancing(transactionHeaderIdx, transaction)

			for i, posting := range transaction.postings {
				validator.checkUnknownAccount(transactionHeaderIdx+i+1, posting)
			}

			transaction = transactionNext
			transactionHeaderIdx = iLine + 1
			continue
		}

		// When the line is a posting, append it to transaction.postings
		postingParseSucceed, posting := parsePostingStr(line)
		if postingParseSucceed {
			transaction.postings = append(transaction.postings, posting)
			continue
		}

		if transaction.date != "" {
			validator.warnPostingParse(transactionHeaderIdx, line)
			continue
		}

		// When the line is neither header or posting, return "Header unmatched" for compatibility
		if transaction.date == "" {
			validator.warnHeaderUnmatched(transactionHeaderIdx)
		}
	}

	validator.checkBalancing(transactionHeaderIdx, transaction)

	for i, posting := range transaction.postings {
		validator.checkUnknownAccount(transactionHeaderIdx+i+1, posting)
	}
}
