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
	}

	for _, fixture := range fixtures {
		caseName := fixture.caseName
		expected := fixture.expectedTransaction
		fixturePath := getFixturePath(caseName)
		bytes, _ := ioutil.ReadFile(fixturePath)
		actualSuceeded, actual := parseTransactionStr(string(bytes))

		if !actualSuceeded {
			t.Errorf("%v: parseTransactionStr failed", fixturePath)
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("%v: %v expected but got %v", fixturePath, expected, actual)
		}
	}
}
