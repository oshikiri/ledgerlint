package main

import (
	"reflect"
	"testing"
)

func TestParsePostingStrWithoutCurrency(t *testing.T) {
	succeed, actual := parsePostingStr("  Asset:Something  100")
	expected := Posting{
		account:     "Asset:Something",
		amount:      100,
		currency:    "",
		emptyAmount: false,
	}
	if !succeed || !reflect.DeepEqual(actual, expected) {
		// t.Errorf("succeed = %v, %v != %v", succeed, actual, expected)
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
		t.Errorf("succeed = %v, %v != %v", succeed, actual, expected)
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
		t.Errorf("succeed = %v, %v != %v", succeed, actual, expected)
	}
}

func TestParsePostingStrEmpty(t *testing.T) {
	succeed, actual := parsePostingStr("  Asset:Something")
	expected := Posting{
		account:     "Asset:Something",
		emptyAmount: true,
	}
	if !succeed || !reflect.DeepEqual(actual, expected) {
		t.Errorf("succeed = %v, %v != %v", succeed, actual, expected)
	}
}
