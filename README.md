# ChatGPT API Example

## Stacks
- Go

## Getting Started
These instructions will help you get a copy of the ChatGPT API Example project up and running on your local machine for development and testing purposes.

### Prerequisites
- Go 1.21 or [later](https://go.dev)

### Installing
1. Clone the project from [ChatGPT API Example Repository](https://github.com/promptsnapshot/chatgpt-api-example)
2. Import the project into your preferred IDE.
3. Copy `config.example.yaml` in the `config` directory, paste it in the same location, and then remove `.example` from its name.
4. Download dependencies by running `go mod download`.

### Testing
1. Run `go test -v -coverpkg ./... -coverprofile coverage.out -covermode count ./...` or use `make test` for testing.

### Running
1. Start the necessary services with `docker-compose up -d` or `make compose-up`.
2. Launch the server by running `go run ./src/.` or use `make server`.

### Handling API Responses
1. Use `make handle-response` to compile and run the script for handling JSON responses from the ChatGPT API.

### Additional Notes
- Ensure you have set up your API keys correctly in the configuration file.
- Familiarize yourself with the ChatGPT API documentation for advanced usage.

### Contributing
Contributions to the ChatGPT API Example project are welcome. Please read [CONTRIBUTING.md](https://github.com/yourusername/chatgpt-api-example/CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.
