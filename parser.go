package main

import (
	"errors"
	"strconv"
)

func consumeWhiteSpace(s string, i int) int {
	for i < len(s) && isWhiteSpace(s[i]) {
		i++
	}
	return i
}

func consumeNonComment(s string, i int) int {
	for i < len(s) && !isCommentSymbol(s[i]) {
		i++
	}
	return i
}

func isDigit(c byte) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}

// TODO: Add currency code if needed
func isCurrencyCode(c byte) bool {
	return c == '$'
}

// TODO: tab?
func isWhiteSpace(c byte) bool {
	return c == ' '
}

func isDateSeparator(c byte) bool {
	return c == '/' || c == '-'
}

func isCommentSymbol(c byte) bool {
	return c == ';'
}

func isStatusSymbol(c byte) bool {
	return c == '!' || c == '*'
}

func isCommentOrEmpty(line string) bool {
	if len(line) == 0 {
		return true
	}

	i := consumeWhiteSpace(line, 0)
	return len(line) == i || line[i] == ';'
}

func parsePostingStr(s string) (bool, Posting) {
	size := len(s)
	succeed := false

	posting := Posting{}
	i := consumeWhiteSpace(s, 0)

	startAccount := i
	for !(isWhiteSpace(s[i]) && isWhiteSpace(s[i+1])) {
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

	if isCurrencyCode(s[i]) {
		posting.currency = string(s[i])
		i++
		amount, _ := strconv.Atoi(s[i:]) // TODO: Error handling
		posting.amount = Amount(amount)
		succeed = true
	} else {
		digitsStart := i
		if s[i] == '-' {
			i++
		}
		// TODO: decimal
		for i < size && isDigit(s[i]) {
			i++
		}
		amount, _ := strconv.Atoi(s[digitsStart:i]) // TODO: Error handling
		posting.amount = Amount(amount)

		if i < size {
			i = consumeWhiteSpace(s, i)
			posting.currency = s[i:]
			succeed = true
		}
	}

	return succeed, posting
}

func parseTransactionHeader(headerIdx int, line string) (Transaction, error) {
	headerUnmatchedError := errors.New("Header unmatched")

	i := consumeWhiteSpace(line, 0)
	if i > 0 {
		return Transaction{}, headerUnmatchedError
	}

	// budger header
	if line[i] == '~' {
		return Transaction{}, nil
	}

	dateStart := i

	t := Transaction{
		postings:  []Posting{},
		headerIdx: headerIdx,
	}

	for i < len(line) && (isDigit(line[i]) || isDateSeparator(line[i])) {
		i++
	}
	t.date = Date(line[dateStart:i])

	iBefore := i
	i = consumeWhiteSpace(line, i)
	if iBefore == i {
		if i == len(line) {
			return t, nil
		} else {
			// it is invalid because non-whitespace character follows date string without whitespace
			return Transaction{}, headerUnmatchedError
		}
	}

	if isStatusSymbol(line[i]) {
		t.status = TransactionStatus(line[i])
		i++
		i = consumeWhiteSpace(line, i)
	}
	startDescription := i
	i = consumeNonComment(line, i)
	t.description = line[startDescription:i]

	return t, nil
}
