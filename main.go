package main

import (
	"flag"
)

func main() {
	var filePath = flag.String("f", "", "ledger/hledger transaction file")
	var accountsPath = flag.String("account", "", "known accounts file")
	var outputJSON = flag.Bool("j", false, "output error message by JSON format or plaintext (default plaintext)")
	flag.Parse()

	if *filePath == "" {
		panic("file path is empty. specify transaction file using '-f' option.")
	}

	lintTransactionFile(
		*filePath,
		*accountsPath,
		*outputJSON,
	)
}
