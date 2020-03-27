ledgerlint
=====

A linter for ledger/hledger transaction files

## Usage
```
$ ledgerlint -f fixtures/imbalance.ledger

2020-03-26 * toilet paper
  Expences:Household essentials    200 JPY
  Assets:Cash                    -2000 JPY
^^^^^
ERROR: imbalanced transaction, total amount = -1800JPY
```

## Installation
### From binary

// TODO: Add goreleaser

### From source code

```sh
go get github.com/oshikiri/ledgerlint
go install github.com/oshikiri/ledgerlint
```

and then add `~/go/bin` to `$PATH`.

## Development
```sh
# Build and create a ./ledgerlint binary
go build

# Run tests
go test
```

## License

MIT
