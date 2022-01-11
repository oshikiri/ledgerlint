package main

import (
	"reflect"
	"testing"
)

func TestParsePostingStrWithoutCurrency(t *testing.T) {
	actual, err := parsePostingStr("  Asset:Something  100")
	expected := Posting{
		account:     "Asset:Something",
		amount:      100,
		currency:    "",
		emptyAmount: false,
	}
	if err != nil || !reflect.DeepEqual(actual, expected) {
		t.Errorf("succeed = %v, %v != %v", err, actual, expected)
	}
}

func TestParsePostingStrJPY(t *testing.T) {
	actual, err := parsePostingStr("  Asset:Something  100 JPY")
	expected := Posting{
		account:     "Asset:Something",
		amount:      100,
		currency:    "JPY",
		emptyAmount: false,
	}
	if err != nil || !reflect.DeepEqual(actual, expected) {
		t.Errorf("err = %v, %v != %v", err, actual, expected)
	}
}

func TestParsePostingStrDollar(t *testing.T) {
	actual, err := parsePostingStr("  Asset:Something  $1")
	expected := Posting{
		account:     "Asset:Something",
		amount:      1,
		currency:    "$",
		emptyAmount: false,
	}
	if err != nil || !reflect.DeepEqual(actual, expected) {
		t.Errorf("err = %v, %v != %v", err, actual, expected)
	}
}

func TestParsePostingStrEmpty(t *testing.T) {
	actual, err := parsePostingStr("  Asset:Something")
	expected := Posting{
		account:     "Asset:Something",
		emptyAmount: true,
	}
	if err != nil || !reflect.DeepEqual(actual, expected) {
		t.Errorf("err = %v, %v != %v", err, actual, expected)
	}
}

func TestParseTransactionHeader(t *testing.T) {
	actual, err := parseTransactionHeader(11, "2020-01-01 * some description")
	expected := Transaction{
		date:        "2020-01-01",
		description: "some description",
		status:      "*",
		postings:    []Posting{},
		headerIdx:   11,
	}
	if err != nil || !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v, %v != %v", err, actual, expected)
	}
}

func TestParseTransactionHeaderInvalid(t *testing.T) {
	actual, err := parseTransactionHeader(11, "2020-01-01* some description")
	expected := Transaction{}
	if err == nil || !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v", err)
	}
}

func TestParsePostingStrComment(t *testing.T) {
	actual, err := parsePostingStr("  Asset:Something  100 ; comment")
	expected := Posting{
		account:     "Asset:Something",
		amount:      100,
		currency:    "",
		emptyAmount: false,
	}
	if err != nil || !reflect.DeepEqual(actual, expected) {
		t.Errorf("succeed = %v, %v != %v", err, actual, expected)
	}
}
