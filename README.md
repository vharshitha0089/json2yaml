# JSON to YAML Converter

A simple Go application that converts JSON data into YAML format using the `github.com/itchyny/json2yaml` library.

## Project Structure

```text
json2yaml/
├── cmd/
│   └── main.go
├── internal/
│   └── converter/
│       └── converter.go
├── go.mod
└── go.sum
```

## Features

- Converts JSON data to YAML format
- Uses the `itchyny/json2yaml` package
- Simple and modular project structure
- Demonstrates package usage in Go

## Installation

Clone the repository:

```bash
git clone <repository-url>
cd json2yaml
```

Install dependencies:

```bash
go get github.com/itchyny/json2yaml
```

## Run the Project

```bash
go run ./cmd
```

## Example

### Input (JSON)

```json
{"name":"Asha Rao","age":28,"active":true}
```

### Output (YAML)

```yaml
name: Asha Rao
age: 28
active: true
```

## How It Works

1. `main.go` provides the JSON input.
2. The `Convert()` function in `converter.go` converts JSON to YAML using the `json2yaml` library.
3. The YAML output is returned and displayed on the console.

## Technologies Used

- Go
- github.com/itchyny/json2yaml

## Author

Harshitha V
