package main

import (
	"fmt"
	"reflect"
	"testing"
)

func getFixturePath(caseName string) string {
	return fmt.Sprintf("fixtures/%v.ledger", caseName)
}

func TestRegexPatternPosting(t *testing.T) {
	postingStr := "  Expenses:Household essentials    200 JPY ; some comments"
	expected := []string{"Expenses:Household essentials", "200", "JPY"}

	actual := postingPattern.FindStringSubmatch(postingStr)[1:]
	if len(actual) != len(expected) || !reflect.DeepEqual(actual, expected) {
		t.Errorf("regex postingPattern should parse as %v but got %v", expected, actual)
	}
}

func TestRegexHeader(t *testing.T) {
	headerStr := "2020-03-26 * toilet paper"
	expected := []string{"2020-03-26", "*", "toilet paper"}
	actual := headerPattern.FindStringSubmatch(headerStr)[1:]
	if len(actual) != len(expected) || !reflect.DeepEqual(actual, expected) {
		t.Errorf("regex headerPattern should parse as %v but got %v", expected, actual)
	}
}

func TestParsePostingStrWithoutCurrency(t *testing.T) {
	succeed, actual := parsePostingStr("  Asset:Something  100")
	expected := Posting{
		account:     "Asset:Something",
		amount:      100,
		currency:    "",
		emptyAmount: false,
	}
	if !succeed || !reflect.DeepEqual(actual, expected) {
		// t.Errorf("%v", actual)
	}
}

func TestParsePostingStrJPY(t *testing.T) {
	succeed, actual := parsePostingStr("  Asset:Something  100 JPY")
	expected := Posting{
		account:     "Asset:Something",
		amount:      100,
		currency:    "JPY",
		emptyAmount: false,
	}
	if !succeed || !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v", actual)
	}
}

func TestParsePostingStrDollar(t *testing.T) {
	succeed, actual := parsePostingStr("  Asset:Something  $1")
	expected := Posting{
		account:     "Asset:Something",
		amount:      1,
		currency:    "$",
		emptyAmount: false,
	}
	if !succeed || !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v", actual)
	}
}

func TestParsePostingStrEmpty(t *testing.T) {
	succeed, actual := parsePostingStr("  Asset:Something")
	expected := Posting{
		account:     "Asset:Something",
		emptyAmount: true,
	}
	if !succeed || !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v", actual)
	}
}
