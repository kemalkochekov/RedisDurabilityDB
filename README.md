# RedisDurabilityDB

This project aims to provide a custom implementation of a caching layer and a storage layer inspired by PostgreSQL and Redis. The implementation is written in the `pkg` package and follows the `datasource.Datasource` interface.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
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

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
