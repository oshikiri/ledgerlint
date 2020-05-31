package main

import (
	"errors"
	"fmt"
	"strings"
)

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
			postingParseError := fmt.Errorf("parsePostingStr is failed: '%v'", line)
			validator.warnParseFailed(transactionHeaderIdx, postingParseError)
			continue
		}

		// When the line is neither header or posting, return "Header unmatched" for compatibility
		if transaction.date == "" {
			err := errors.New("Header unmatched")
			validator.warnParseFailed(transactionHeaderIdx, err)
		}
	}

	validator.checkBalancing(transactionHeaderIdx, transaction)

	for i, posting := range transaction.postings {
		validator.checkUnknownAccount(transactionHeaderIdx+i+1, posting)
	}
}
