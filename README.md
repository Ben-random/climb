# climb 🧗

CLI tool that makes your local scripts globally available

## Usage 

```bash
./path/to/script
> Hello

climb create example path/to/script

example
> Hello
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
