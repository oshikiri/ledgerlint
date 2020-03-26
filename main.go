package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var filePath = flag.String("f", "", "ledger/hledger transaction file")
	flag.Parse()

	fmt.Println("file path =", *filePath)

	bytes, err := ioutil.ReadFile(*filePath)
	if err != nil {
		panic(err)
	}

	transactionStrs := strings.Split(string(bytes), "\n\n")
	fmt.Println(parseTransactionStr(transactionStrs[0]))
}
