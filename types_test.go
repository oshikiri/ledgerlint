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

func TestTransactionEmptyAmount(t *testing.T) {
	tx := transactionsBalancedEmptyAmount
	actual := tx.calculateTotalAmount()
	expected := map[string]Amount{
		"JPY": 800,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%v != %v", actual, expected)
	}
}
