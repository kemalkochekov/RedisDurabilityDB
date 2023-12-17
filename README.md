# RedisDurabilityDB

This project aims to provide a custom implementation of a caching layer and a storage layer inspired by PostgreSQL and Redis. The implementation is written in the `pkg` package and follows the `datasource.Datasource` interface.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Linting and Code Quality](#linting-and-code-quality)
  - [Linting Installation](#linting-installation)
  - [Linting Usage](#linting-usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- GET: Retrieve data from the database based on the provided key.
- INSERT: Add new data to the database.

## Prerequisites

Before running this application, ensure that you have the following prerequisites installed:

- Go: [Install Go](https://go.dev/doc/install/)

## Installation

1. Clone the repository:
  ```bash
    https://github.com/kemalkochekov/RedisDurabilityDB.git
  ```

2. Navigate to the project directory:
  ```bash
    cd RedisDurabilityDB
  ```

## Usage
1. Run the main.go file:
  ```bash
    go run cmd/main.go
  ```

## Linting and Code Quality

This project maintains code quality using `golangci-lint`, a fast and customizable Go linter. `golangci-lint` checks for various issues, ensures code consistency, and enforces best practices, helping maintain a clean and standardized codebase.

### Linting Installation

To install `golangci-lint`, you can use `brew`:

```bash
  brew install golangci-lint
```

### Linting Usage

1. Configuration: 

After installing golangci-lint, create or use a personal configuration file (e.g., .golangci.yml) to define specific linting rules and settings:
```bash
  golangci-lint run --config=.golangci.yml
```
This command initializes linting based on the specified configuration file.

2. Run the linter:

Once configuration is completed, you can execute the following command at the root directory of your project to run golangci-lint:

```bash
  golangci-lint run
```
This command performs linting checks on your entire project and provides a detailed report highlighting any issues or violations found.

3. Customize Linting Rules:

You can customize the linting rules by modifying the `.golangci.yml` file.

For more information on using golangci-lint, refer to the golangci-lint documentation.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
