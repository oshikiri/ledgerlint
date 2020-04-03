package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var commentOrEmptyPattern = regexp.MustCompile(`^\s*(?:;|$)`)
var headerPattern = regexp.MustCompile(`^(~|\d{4}[-\/]\d{2}[-\/]\d{2})(?:\s+(?:([\*!])\s+|)([^;]+))?(?:;.+)?$`)
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
	return false, Posting{}
}

func parsePostingStrs(postingStrs []string) ([]Posting, error) {
	postings := []Posting{}

	for _, postingStr := range postingStrs {
		if commentOrEmptyPattern.MatchString(postingStr) {
			continue
		}
		succeed, posting := parsePostingStr(postingStr)
		if succeed {
			postings = append(postings, posting)
		} else {
			msg := fmt.Sprintf("parsePostingStr is failed: '%v'", postingStr)
			return nil, errors.New(msg)
		}
	}

	return postings, nil
}

func parseTransactionStr(s string) (Transaction, error) {
	lines := strings.Split(s, "\n")
	i := 0

	// Skip comment or empty lines
	for _, line := range lines {
		if !commentOrEmptyPattern.MatchString(line) {
			break
		}
		i++
	}

	matched := headerPattern.FindStringSubmatch(lines[i])
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

	postings, err := parsePostingStrs(lines[(i + 1):])
	if err != nil {
		return Transaction{}, err
	}
	t.postings = postings
	return t, nil
}
