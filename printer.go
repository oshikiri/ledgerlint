package main

import (
	"fmt"
	"strings"
)

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

func (printer *Printer) warnParseFailed(countNewlines int) {
	printer.print(
		countNewlines,
		"WARN",
		fmt.Errorf("This line is neither comment nor header nor posting"),
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
	severity := 3
	if logLevel == "ERROR" {
		severity = 1
	} else if logLevel == "WARN" {
		severity = 2
	}
	if printer.outputJSON {
		errorMessage := strings.ReplaceAll(err.Error(), "\"", "\\\"")
		parseFailedMsg := `{"type":"diagnostic","source":"ledgerlint","file_path":"%v","line_number":%v,"range":{"start":{"line":%v,"character":%v},"end":{"line":%v,"character":%v}},"level":"%v","severity":%v,"message":"%v"}` + "\n"
		fmt.Printf(parseFailedMsg, printer.filePath, countNewlines, countNewlines-1, 0, countNewlines-1, 80, logLevel, severity, errorMessage)
	} else {
		parseFailedMsg := "%v:%v %v\n"
		fmt.Printf(parseFailedMsg, printer.filePath, countNewlines, err)
	}
}
