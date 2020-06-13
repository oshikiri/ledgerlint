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
