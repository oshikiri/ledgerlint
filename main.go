package main

import (
	"flag"
)

func main() {
	var filePath = flag.String("f", "", "ledger/hledger transaction file")
	var accountsPath = flag.String("account", "", "known accounts file")
	flag.Parse()

	lintTransactionFile(*filePath, *accountsPath)
}
