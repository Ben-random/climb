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
./bin/pkg <command> <alias> <path/to/script NOTE: only for create or update>
```

or
```bash
go run main.go <command> <alias> <path/to/script NOTE: only for create or update>
```
