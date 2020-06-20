package main

import (
	"errors"
	"regexp"
	"strconv"
)

// FIXME: dirty parsing logic https://github.com/oshikiri/ledgerlint/issues/18
var commentOrEmptyPattern = regexp.MustCompile(`^\s*(?:;|$)`)
var headerPattern = regexp.MustCompile(`^(~|\d{4}[-\/]\d{2}[-\/]\d{2})(?:\s+(?:([\*!])\s+|)([^;]+))?(?:;.+)?$`)

func consumeWhiteSpace(s string, i int) int {
	for s[i] == ' ' {
		i++
	}
	return i
}

func isDigit(c byte) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}

func parsePostingStr(s string) (bool, Posting) {
	size := len(s)

	posting := Posting{}
	i := consumeWhiteSpace(s, 0)

	startAccount := i
	for !(s[i] == ' ' && s[i+1] == ' ') {
		i++
		if i >= size {
			break
		}
	}
	posting.account = s[startAccount:i]
	if i == size {
		posting.emptyAmount = true
		return true, posting
	}

	posting.emptyAmount = false
	i += 2

	if s[i] == '$' {
		posting.currency = "$"
		i++
		amount, _ := strconv.Atoi(s[i:])
		posting.amount = Amount(amount)
	} else {
		j := i
		for j < size && isDigit(s[j]) {
			j++
		}
		amount, _ := strconv.Atoi(s[i:j])
		posting.amount = Amount(amount)
		i = j

		if i < size {
			i = consumeWhiteSpace(s, i)
			posting.currency = s[i:]
		}
	}

	return true, posting
}

func parseTransactionHeader(headerIdx int, line string) (Transaction, error) {
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
		headerIdx:   headerIdx,
	}

	return t, nil
}
