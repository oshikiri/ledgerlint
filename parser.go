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

func consumeNonWhiteSpace(s string, i int) int {
	for i < len(s) && !isWhiteSpace(s[i]) {
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

func consumeDateStr(s string, i int) int {
	for i < len(s) && (isDigit(s[i]) || isDateSeparator(s[i])) {
		i++
	}
	return i
}

func consumeDigits(s string, i int) int {
	for i < len(s) && isDigit(s[i]) {
		i++
	}
	return i
}

func consumeUntilDoubleWhiteSpace(s string, i int) int {
	for i < len(s) && !(isWhiteSpace(s[i]) && isWhiteSpace(s[i+1])) {
		i++
	}
	return i
}

func restoreTailWhiteSpaces(s string, i int) int {
	if s[i-1] != ' ' {
		return i
	}
	i--
	for i >= 0 && isWhiteSpace(s[i]) {
		i--
	}
	return i + 1
}

func isDigit(c byte) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}

// TODO: Add currency code if needed
func isCurrencyCode(c byte) bool {
	return c == '$'
}

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
	return len(line) == i || isCommentSymbol(line[i])
}

func parseAmount(s string) (Amount, error) {
	amount, err := strconv.Atoi(s)
	if err == nil {
		return Amount(amount), nil
	}
	return -1, errors.New("Invalid amount")
}

func parsePostingStr(s string) (Posting, error) {
	size := len(s)
	posting := Posting{}
	var err error

	i := consumeWhiteSpace(s, 0)

	if i == 0 {
		return posting, errors.New("Posting without indents")
	}

	startAccount := i
	i = consumeUntilDoubleWhiteSpace(s, i)
	posting.account = s[startAccount:i]
	if i == size {
		posting.emptyAmount = true
		return posting, nil
	}

	i = consumeWhiteSpace(s, i)

	var amount Amount
	var currency string
	i, amount, currency, err = parseAccountAndCurrency(s, i)
	i = consumeWhiteSpace(s, i)
	if i+1 < size && s[i] == '@' && s[i+1] == '@' {
		i = i + 2
		i = consumeWhiteSpace(s, i)
		i, amount, currency, err = parseAccountAndCurrency(s, i)
	}

	posting.amount = amount
	posting.currency = currency

	return posting, err
}

func parseAccountAndCurrency(s string, i int) (int, Amount, string, error) {
	currency := ""
	var amount Amount
	var err error
	size := len(s)

	if isCurrencyCode(s[i]) {
		currency = string(s[i])
		i++
		amount, err = parseAmount(s[i:])
	} else {
		digitsStart := i
		if s[i] == '-' {
			i++
		}
		// TODO: decimal
		i = consumeDigits(s, i)
		amount, err = parseAmount(s[digitsStart:i])

		if i < size {
			i = consumeWhiteSpace(s, i)
			start := i
			i = consumeNonWhiteSpace(s, i)
			i = restoreTailWhiteSpaces(s, i)
			currency = s[start:i]
		}
	}

	return i, amount, currency, err
}

func parseTransactionHeader(headerIdx int, line string) (Transaction, error) {
	size := len(line)

	i := consumeWhiteSpace(line, 0)
	if i > 0 {
		return Transaction{}, errors.New("Non-header")
	}

	// budget header
	if line[i] == '~' {
		i++
		i = consumeWhiteSpace(line, i)
	}
	if i+5 < size && line[i:(i+5)] == "every" {
		return Transaction{description: line[i:]}, nil
	}

	dateStart := i

	t := Transaction{
		postings:  []Posting{},
		headerIdx: headerIdx,
	}

	i = consumeDateStr(line, i)
	t.date = Date(line[dateStart:i])

	iBefore := i
	i = consumeWhiteSpace(line, i)
	if iBefore == i {
		if i == len(line) {
			return t, nil
		}
		return Transaction{}, errors.New("Non-whitespace character follows date string without whitespace")
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
