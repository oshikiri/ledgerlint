package main

import (
	"bufio"
	"fmt"
	"os"
)

func lintTransactionFile(filePath, accountsPath string, outputJSON bool) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	transaction := Transaction{headerIdx: 1}
	validator := newValidator(filePath, accountsPath, outputJSON)

	for iLine := 1; scanner.Scan(); iLine++ {
		if err := scanner.Err(); err != nil {
			fmt.Printf("Scanner failed: %v\n", err)
			break
		}
		line := scanner.Text()

		// When the line is empty, skip it
		if isCommentOrEmpty(line) {
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
		posting, postingParseError := parsePostingStr(line)

		if postingParseError == nil {
			if posting.emptyAmount && transaction.date == "" {
				validator.printer.warnHeaderUnmatched(transaction.headerIdx)
			} else {
				transaction.postings = append(transaction.postings, posting)
			}
			continue
		}

		// When the line is market price, skip it for now
		// https://hledger.org/investments.html#market-prices
		if line[0] == 'P' {
			continue
		}

		if transaction.date == "" {
			validator.printer.warnParseFailed(iLine)
		} else {
			validator.printer.warnPostingParse(transaction.headerIdx, line)
		}
	}

	validator.checkBalancingAndAccounts(transaction)
}
