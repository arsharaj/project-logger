## Project : Logger

> Module to parse system logs

### Features

- Configuration Management
- Log File Tailing
- Parsing Logs
- Elasticsearch Integration
- Unit Testing

### Requirements

- go : 1.24.4

### Usage

1. Run application : `go run ./cmd`
2. Clean go dependencies : `go mod tidy`

### Folder Structure

```
root
├── .github
├── cmd
├── config
├── docs
├── elk
├── internal
├── parser
├── tailer
└── test
```

### Documentation

Full documentation can be found in [docs](docs/) folder.

### Running Tests

Run the following command to execute the test suite : `go test ./test`

### Contributing

We welcome contributions! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Versioning

This project uses [Semantic Versioning](https://semver.org/).

### License

[MIT License](LICENSE)