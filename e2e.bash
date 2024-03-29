count_failed=0

function test_ledgerlint() {
  spec_case=$1
  args=$2

  echo "CASE ${spec_case}:"
  diff "fixtures/${spec_case}.expected.txt" <(./ledgerlint -f fixtures/${spec_case}.ledger ${args})
  if [ $? -eq 0 ]; then
    echo -e "\tPASSED plaintext"
  else
    count_failed=$(($count_failed + 1))
    echo -e "\tFAILED plaintext"
  fi

  diff "fixtures/${spec_case}.expected.json" <(./ledgerlint -f fixtures/${spec_case}.ledger -j ${args})
  if [ $? -eq 0 ]; then
    echo -e "\tPASSED json"
  else
    count_failed=$(($count_failed + 1))
    echo -e "\tFAILED json"
  fi
  echo -e ""
}

function pending {
  command=$0
  echo -e "CASE $2:\n\tPENDING\n"
}

test_ledgerlint balanced
test_ledgerlint balanced-empty-amount
test_ledgerlint imbalance
test_ledgerlint imbalance-multi-currency
test_ledgerlint unknown-account "-account fixtures/accounts.txt"
test_ledgerlint unmatched
test_ledgerlint no-description
test_ledgerlint nonewline
test_ledgerlint budget
test_ledgerlint two-empty-amount
test_ledgerlint dollar-mark
test_ledgerlint comment
test_ledgerlint unparsable

if [ $count_failed -eq 0 ]; then
  echo "All tests passed"
  exit 0
else
  echo "${count_failed} failed cases"
  exit 1
fi
