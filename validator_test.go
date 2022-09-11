package main

func ExampleValidatorHeaderUnmatched() {
	validator := newValidator("example/some.ledger", "", false)
	validator.printer.warnHeaderUnmatched(11)
	// Output: example/some.ledger:11 Header unmatched
}

func ExampleValidatorHeaderUnmatchedJSON() {
	validator := newValidator("example/some.ledger", "", true)
	validator.printer.warnHeaderUnmatched(11)
	// Output: {"type":"diagnostic","source":"ledgerlint","file_path":"example/some.ledger","line_number":11,"range":{"start":{"line":10,"character":0},"end":{"line":10,"character":80}},"level":"WARN","severity":2,"message":"Header unmatched"}
}
