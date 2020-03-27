package main

import (
	"regexp"
	"strconv"
	"strings"
)

var headerPattern = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}) (\*) (.+)`)
var postingPattern = regexp.MustCompile(`\s{2,}(.+)\s{2,}(-?\s?\d+)\s(JPY)`)

func parsePostingStr(s string) (bool, Posting) {
	m := postingPattern.FindStringSubmatch(s)
	if len(m) < 4 {
		return false, Posting{}
	}
	figure, _ := strconv.Atoi(m[2])
	p := Posting{
		account:  m[1],
		amount:   figure,
		currency: m[3],
	}

	return true, p
}

func parseTransactionStr(s string) Transaction {
	lines := strings.Split(s, "\n")
	header := headerPattern.FindAllStringSubmatch(lines[0], 1)[0][1:]
	t := Transaction{
		date:        header[0],
		status:      header[1],
		description: header[2],
		postings:    []Posting{},
	}
	postingStrs := lines[1:]
	for _, postingStr := range postingStrs {
		succeed, posting := parsePostingStr(postingStr)
		if succeed {
			t.postings = append(t.postings, posting)
		}
	}
	return t
}
