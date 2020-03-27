package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func getFixturePath(caseName string) string {
	return fmt.Sprintf("fixtures/%v.ledger", caseName)
}

func TestRegexPatternPosting(t *testing.T) {
	postingStr := "  Expences:Household essentials  200 JPY ; some comments"
	expected := []string{"Expences:Household essentials", "200", "JPY"}

	actual := postingPattern.FindStringSubmatch(postingStr)[1:]
	if len(actual) != len(expected) || !reflect.DeepEqual(actual, expected) {
		t.Errorf("regex patternPosting should parse as %v but got %v", expected, actual)
	}
}

func TestReadFixtureImbalance(t *testing.T) {
	caseName := "imbalance"
	fixturePath := getFixturePath(caseName)

	bytes, _ := ioutil.ReadFile(fixturePath)
	actual := parseTransactionStr(string(bytes))
	expected := transactionsImbalanced

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Read %v, %v expected but got %v", fixturePath, expected, actual)
	}
}

func TestReadFixtureImbalanceMultiCurrency(t *testing.T) {
	caseName := "imbalance-multi-currency"
	fixturePath := getFixturePath(caseName)

	bytes, _ := ioutil.ReadFile(fixturePath)
	actual := parseTransactionStr(string(bytes))
	expected := transactionsImbalancedMultiCurrency

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Read %v, %v expected but got %v", fixturePath, expected, actual)
	}
}
