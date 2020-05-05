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

type FixtureTuple struct {
	caseName            string
	expectedTransaction Transaction
}

func TestReadFixtures(t *testing.T) {
	fixtures := []FixtureTuple{
		{
			"imbalance",
			transactionsImbalanced,
		},
		{
			"imbalance-multi-currency",
			transactionsImbalancedMultiCurrency,
		},
		{
			"balanced",
			transactionsBalanced,
		},
		{
			"balanced-empty-amount",
			transactionsBalancedEmptyAmount,
		},
	}

	for _, fixture := range fixtures {
		caseName := fixture.caseName
		expected := fixture.expectedTransaction
		fixturePath := getFixturePath(caseName)
		bytes, _ := ioutil.ReadFile(fixturePath)
		actual, err := parseTransactionStr(string(bytes))

		if err != nil {
			t.Errorf("%v: parseTransactionStr failed '%v'", fixturePath, err)
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v: %v expected but got %v", fixturePath, expected, actual)
		}
	}
}
