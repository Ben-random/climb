# climb 🧗

CLI tool that makes your local scripts and binaries able to be called from anywhere in your terminal.

## Example Usage 

```bash
./path/to/script
> Hello

climb create example path/to/script

example
> Hello
```

## Usage

```
Usage: climb <command> <alias> [script-path]

Commands:
  create <alias> <script-path>  Create a new alias for a script
  update <alias> <script-path>  Update an existing alias
  delete <alias>                Delete an existing alias
  help                          Show this help message

Options:
  --dry-run                     Preview changes without modifying files
```

## Documentation

See docs [here](https://github.com/Ben-random/climb/tree/main/docs/index.md)

## Installation

To install this cli, you will first need the `go` package if you haven't already. You can find instructions to install it [here](https://go.dev/doc/install).

run:
```bash
chmod +x ./scripts/install.sh
./scripts/install.sh
```

## Development

Build

```bash
go build -o ./bin/climb
```

run
```bash
./bin/pkg create myAlias ./path/to/myScript
```

or
```bash
go run main.go create myAlias ./path/to/myScript
```
