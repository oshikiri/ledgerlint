package main

import "strings"

func lintTransactionFile(filePath, accountsPath string) {
	validator := newValidator(filePath, accountsPath)

	countNewlines := 1
	transactionsStr, err := readFileContent(filePath)
	if err != nil {
		panic(err)
	}

	for _, transactionStr := range strings.Split(transactionsStr, "\n\n") {
		success, transaction := parseTransactionStr(transactionStr)
		if success {
			validator.checkBalancing(countNewlines, transaction)

			for i, posting := range transaction.postings {
				validator.checkUnknownAccount(countNewlines+i+1, posting)
			}
		} else {
			validator.warnParseFailed(countNewlines)
		}

		countNewlines += strings.Count(transactionStr, "\n") + 2
	}
}
