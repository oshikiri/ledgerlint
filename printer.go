package main

import "fmt"

// Printer emits warning/error messages
type Printer struct {
	filePath   string
	outputJSON bool
}

func (printer *Printer) warnUnknownAccount(countNewlines int, account string) {
	printer.print(
		countNewlines,
		"WARN",
		fmt.Errorf("unknown account: %v", account),
	)
}

func (printer *Printer) warnHeaderUnmatched(countNewlines int) {
	printer.print(
		countNewlines,
		"WARN",
		fmt.Errorf("Header unmatched"),
	)
}

func (printer *Printer) warnPostingParse(countNewlines int, line string) {
	printer.print(
		countNewlines,
		"WARN",
		fmt.Errorf("parsePostingStr is failed: '%v'", line),
	)
}

func (printer *Printer) print(countNewlines int, logLevel string, err error) {
	if printer.outputJSON {
		parseFailedMsg := `{"file_path":"%v","line_number":%v,"level":"%v","error_message":"%v"}` + "\n"
		fmt.Printf(parseFailedMsg, printer.filePath, countNewlines, logLevel, err)
	} else {
		parseFailedMsg := "%v:%v %v\n"
		fmt.Printf(parseFailedMsg, printer.filePath, countNewlines, err)
	}
}
