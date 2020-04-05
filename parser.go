package main

import (
	"errors"
	"regexp"
	"strconv"
)

// FIXME: suspicious parsing logic
var commentOrEmptyPattern = regexp.MustCompile(`^\s*(?:;|$)`)
var headerPattern = regexp.MustCompile(`^(~|\d{4}[-\/]\d{2}[-\/]\d{2})(?:\s+(?:([\*!])\s+|)([^;]+))?(?:;.+)?$`)
var postingPattern = regexp.MustCompile(`\s{2,}([^;]+\S)\s{2,}(-?\s?\d+)\s([\w^;]+)`)
var postingPatternWithCurrencyMark = regexp.MustCompile(`\s{2,}([^;]+\S)\s{2,}(\$)(-?\s?\d+)`)
var postingEmptyAmountPattern = regexp.MustCompile(`\s{2,}([^;]+)`)

func parsePostingStr(s string) (bool, Posting) {
	m := postingPattern.FindStringSubmatch(s)
	if len(m) == 4 { // non-empty amount
		amount, err := strconv.Atoi(m[2])
		if err == nil {
			p := Posting{
				account:     m[1],
				amount:      Amount(amount),
				currency:    m[3],
				emptyAmount: false,
			}
			return true, p
		}
	}

	m = postingPatternWithCurrencyMark.FindStringSubmatch(s)
	if len(m) == 4 {
		amount, err := strconv.Atoi(m[3])
		if err == nil {
			p := Posting{
				account:     m[1],
				amount:      Amount(amount),
				currency:    m[2],
				emptyAmount: false,
			}
			return true, p
		}
	}

	if len(m) == 0 { // empty amount
		m := postingEmptyAmountPattern.FindStringSubmatch(s)
		if len(m) == 2 {
			p := Posting{
				account:     m[1],
				emptyAmount: true,
			}
			return true, p
		}
	}

	return false, Posting{}
}

func parseTransactionHeader(line string) (Transaction, error) {
	matched := headerPattern.FindStringSubmatch(line)
	if len(matched) == 0 {
		return Transaction{}, errors.New("Header unmatched")
	}

	header := matched[1:]
	t := Transaction{
		date:        Date(header[0]),
		status:      TransactionStatus(header[1]),
		description: header[2],
		postings:    []Posting{},
	}

	return t, nil
}
