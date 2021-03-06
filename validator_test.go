package main

func ExampleValidatorHeaderUnmatched() {
	validator := newValidator("example/some.ledger", "", false)
	validator.printer.warnHeaderUnmatched(11)
	// Output: example/some.ledger:11 Header unmatched
}

func ExampleValidatorHeaderUnmatchedJSON() {
	validator := newValidator("example/some.ledger", "", true)
	validator.printer.warnHeaderUnmatched(11)
	// Output: {"file_path":"example/some.ledger","line_number":11,"level":"WARN","error_message":"Header unmatched"}
}
