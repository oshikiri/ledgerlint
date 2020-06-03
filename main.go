package main

import (
	"flag"
	"fmt"
)

var version = "unspecified" // filled by ldflags option

func main() {
	var filePath = flag.String("f", "", "ledger/hledger transaction file")
	var accountsPath = flag.String("account", "", "known accounts file")
	var outputJSON = flag.Bool("j", false, "output error message by JSON format or plaintext (default plaintext)")
	var showVersion = flag.Bool("v", false, "show version and exit")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	if *filePath == "" {
		panic("file path is empty. specify transaction file using '-f' option.")
	}

	lintTransactionFile(
		*filePath,
		*accountsPath,
		*outputJSON,
	)
}
