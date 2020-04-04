ledgerlint: a linter for ledger transaction
=====

[![Build Status](https://github.com/oshikiri/ledgerlint/workflows/Go/badge.svg)](https://github.com/oshikiri/ledgerlint/actions?query=workflow%3AGo) [![go report](https://goreportcard.com/badge/github.com/oshikiri/ledgerlint)](https://goreportcard.com/report/github.com/oshikiri/ledgerlint) [![release](https://img.shields.io/github/v/release/oshikiri/ledgerlint.svg)](https://github.com/oshikiri/ledgerlint/releases/latest)

```sh
$ cat fixtures/imbalance.ledger
2020-03-26 * toilet paper
  Expences:Household essentials    200 JPY
  Assets:Cash                    -2000 JPY

$ ledgerlint -f fixtures/imbalance.ledger
fixtures/imbalance.ledger:1 imbalanced transaction, (total amount) = -1800 JPY
```

If you use vscode, see [vscode-ledgerlint].

[vscode-ledgerlint]: https://github.com/oshikiri/vscode-ledgerlint

## Installation

```sh
./docs/install.bash
```

See also <https://github.com/oshikiri/ledgerlint/releases/latest>

## Development
```sh
# Build and create a ./ledgerlint binary
go build

# Run tests
go test
```

Install from source code:

```sh
go get github.com/oshikiri/ledgerlint
go install github.com/oshikiri/ledgerlint
```

and then add `~/go/bin` to `$PATH`.

## License

MIT
