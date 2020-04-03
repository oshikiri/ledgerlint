package main

import (
	"reflect"
	"testing"
)

type TransactionFixtureTuple struct {
	transaction                    Transaction
	expectedContainsOneEmptyAmount bool
	expectedAmounts                map[string]Amount
}

func TestTransactionCalculateTotalAmount(t *testing.T) {
	fixtures := []TransactionFixtureTuple{
		{
			transactionsBalancedEmptyAmount,
			true,
			map[string]Amount{
				"JPY": 800,
			},
		},
		{
			transactionsImbalanced,
			false,
			map[string]Amount{
				"JPY": -1800,
			},
		},
	}

	for _, fixture := range fixtures {
		transaction := fixture.transaction
		expectedContainsOneEmptyAmount := fixture.expectedContainsOneEmptyAmount
		expectedAmounts := fixture.expectedAmounts

		actualContainsOneEmptyAmount, actualTotalAmounts, err := transaction.calculateTotalAmount()
		if err != nil {
			t.Errorf("error is not nil: %v", err)
		}
		if actualContainsOneEmptyAmount != expectedContainsOneEmptyAmount {
			t.Errorf("containsOneEmptyAmount should be %v but got %v", expectedContainsOneEmptyAmount, actualContainsOneEmptyAmount)
		}
		if !reflect.DeepEqual(actualTotalAmounts, expectedAmounts) {
			t.Errorf("%v != %v", actualTotalAmounts, expectedAmounts)
		}
	}
}
