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
