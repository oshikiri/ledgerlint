package main

func ExampleValidatorHeaderUnmatched() {
	validator := newValidator("example/some.ledger", "", false)
	validator.printer.warnHeaderUnmatched(11)
	// Output: example/some.ledger:11 Header unmatched
}

func ExampleValidatorHeaderUnmatchedJSON() {
	validator := newValidator("example/some.ledger", "", true)
	validator.printer.warnHeaderUnmatched(11)
	// Output: {"source":"ledgerlint","file_path":"example/some.ledger","line_number":11,"range":{"startLine":10,"startCharacter":0,"endLine":10,"endCharacter":80},"level":"WARN","severity":2,"message":"Header unmatched"}
}
