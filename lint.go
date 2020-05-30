package main

import "strings"

func lintTransactionFile(filePath, accountsPath string) {
	validator := newValidator(filePath, accountsPath)

	countNewlines := 1
	transactionsStr, err := readFileContent(filePath)
	if err != nil {
		panic(err)
	}

	transactionStrs := strings.Split(transactionsStr, "\n\n")
	for _, transactionStr := range transactionStrs {
		countNewlinesOld := countNewlines
		countNewlines += strings.Count(transactionStr, "\n") + 2

		transaction, err := parseTransactionStr(transactionStr)
		if err != nil {
			validator.warnParseFailed(countNewlinesOld, err)
			continue
		}

		validator.checkBalancing(countNewlinesOld, transaction)

		for i, posting := range transaction.postings {
			validator.checkUnknownAccount(countNewlinesOld+i+1, posting)
		}
	}
}
