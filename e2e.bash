count_failed=0
spec_cases=("balanced" "balanced-empty-amount" "imbalance" "unknown-account")

for spec_case in "${spec_cases[@]}"
do
  echo "CASE ${spec_case}:"
  diff "fixtures/${spec_case}-output.txt" <(./ledgerlint -f fixtures/${spec_case}.ledger)
  if [ $? -eq 0 ]; then
    echo -e "\tPASSED"
  else
    count_failed=$(($count_failed + 1))
    echo -e "\tFAILED"
  fi
  echo -e ""
done

if [ $count_failed -eq 0 ]; then
  echo "All tests passed"
  exit 0
else
  echo "${count_failed} failed cases"
  exit 1
fi
