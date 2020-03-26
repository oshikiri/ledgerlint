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

func TestReadFixtureImbalance(t *testing.T) {
	caseName := "imbalance"
	fixturePath := getFixturePath(caseName)

	bytes, _ := ioutil.ReadFile(fixturePath)
	actual := parseTransactionStr(string(bytes))
	expected := Transaction{
		date:        "2020-03-26",
		status:      "*",
		description: "toilet paper",
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Read %v, %v expected but got %v", fixturePath, expected, actual)
	}
}
