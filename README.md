ledgerlint
=====

A linter for ledger/hledger transaction files

## Usage
```
$ ledgerlint -f fixtures/imbalance.ledger

# TODO
```

## Installation
### From binary

// TODO

## From source code

```sh
go get github.com/oshikiri/scount
go install github.com/oshikiri/scount
```

and then add `~/go/bin` to `$PATH`.

## Development
```sh
# Build and create ./ledgerlint binary
go build

# Run tests
go test
```

## License

MIT

## TODOs
- lint
    - imbalance amounts
    - unknown accounts
    - invalid spaces
    - aesthetic spacing
- Read ledger/hledger parsing specification/implementation
- `--overwrite`
- Impl goreleaser
- Add vscode extension (vscode-ledgerlint)
- config file
