package main

import (
	"reflect"
	"testing"
)

func TestTransactionTotalAmount(t *testing.T) {
	tx := transactionsImbalanced
	actual := tx.calculateTotalAmount()
	expected := map[string]Amount{
		"JPY": -1800,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}
