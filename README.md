ledgerlint: a linter for ledger transaction
=====

```sh
$ cat fixtures/imbalance.ledger
2020-03-26 * toilet paper
  Expences:Household essentials    200 JPY
  Assets:Cash                    -2000 JPY

$ ledgerlint -f fixtures/imbalance.ledger
fixtures/imbalance.ledger:1 imbalanced transaction, (total amount) = -1800 JPY
```

ledgerlint checks

- imbalanced transactions
- unknown account names

See `ledgerlint -h` for details.

If you use vscode, see [vscode-ledgerlint].

[vscode-ledgerlint]: https://github.com/oshikiri/vscode-ledgerlint

## Installation
### From binary

See <https://github.com/oshikiri/ledgerlint/releases>

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
