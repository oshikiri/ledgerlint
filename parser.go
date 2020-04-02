package main

import (
	"regexp"
	"strconv"
	"strings"
)

var headerPattern = regexp.MustCompile(`(\d{4}[-\/]\d{2}[-\/]\d{2})(?:\s+(?:([\*!])\s+|)([^;]+))?(?:;.+)?$`)
var postingPattern = regexp.MustCompile(`\s{2,}([^;]+)\s{2,}(-?\s?\d+)\s([\w^;]+)`)
var postingEmptyAmountPattern = regexp.MustCompile(`\s{2,}([^;]+)`)

func parsePostingStr(s string) (bool, Posting) {
	m := postingPattern.FindStringSubmatch(s)
	if len(m) == 0 { // empty amount
		m := postingEmptyAmountPattern.FindStringSubmatch(s)
		if len(m) == 2 {
			p := Posting{
				account:     m[1],
				emptyAmount: true,
			}
			return true, p
		}
	} else if len(m) == 4 { // non-empty amount
		amount, _ := strconv.Atoi(m[2]) // FIXME: error handling
		p := Posting{
			account:     m[1],
			amount:      Amount(amount),
			currency:    m[3],
			emptyAmount: false,
		}
		return true, p
	}
	return false, Posting{}
}

func parseTransactionStr(s string) (bool, Transaction) {
	lines := strings.Split(s, "\n")
	matched := headerPattern.FindStringSubmatch(lines[0])
	if len(matched) == 0 {
		return false, Transaction{}
	}

	header := matched[1:]
	t := Transaction{
		date:        Date(header[0]),
		status:      TransactionStatus(header[1]),
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
	return true, t
}
