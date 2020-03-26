package main

import (
	"regexp"
	"strings"
)

var headerPattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}) (\*) (.+)`)

func parseTransactionStr(s string) Transaction {
	lines := strings.Split(s, "\n")
	header := headerPattern.FindAllStringSubmatch(lines[0], 1)[0][1:]
	t := Transaction{
		date:        header[0],
		status:      header[1],
		description: header[2],
	}
	return t
}
