# climb 🧗

CLI tool that makes your local scripts globally available

## Example Usage 

```bash
./path/to/script
> Hello

climb create example path/to/script

example
> Hello
```

## Documentation

See docs [here](https://github.com/Ben-random/climb/tree/main/docs)

## Installation

To install this cli, you will first need the `go` package if you haven't already. You can find instructions to install it [here](https://go.dev/doc/install).

1.) First - You'll need to build the executable:

```bash
go build -o ./bin/climb
```
2.) Next - you need to run the install script

run:
```bash
chmod +x ./scripts/install.sh
./scripts/install.sh
```

## Development

Build

```bash
go build -o bin/pkg
```

run
```bash
./bin/pkg create myAlias ./path/to/myScript
```

or
```bash
go run main.go create myAlias ./path/to/myScript
```
