
# Backend
This repository contains the code for the Hallalins API written in Go. It provides a starting point for building a Go service with common features such as migrations, changelog generation, linting, testing, and more.

## Prerequisites

Before running the commands in this repository, make sure you have the following installed:

- Go (version 1.21.5)
- Git
- [git-chglog](https://github.com/git-chglog/git-chglog)
- [golangci-lint 1.52.2](https://golangci-lint.run)
- All dependencies from config.yml

## Installation

Clone the repository and navigate to the project directory:

```bash
git clone git@github.com:Halalins/backend.git
cd boilerplate-service
```

## Configuration
Copy the configuration file template:
```bash
cp config.yml.example config.yml
```

Open the `config.yml` file in a text editor and fill in the necessary values based on your environment and configuration requirements. Save the file when you're done.

## Usage
### Generate RSA Key
To run the RSA key generator, use the following command:
```
openssl genrsa -out private.pem 2048
openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```
And to get base64 format, use the following command:
```
base64 ./private.pem
base64 ./public.pem
```

### Running the Server
To run the server, use the following command:
```bash
make run
```

This command starts the boilerplate service server.

To run the server with live-reloading (during development only), use the following command:
```bash
make run_server_dev
```
This command starts the boilerplate service server that reload everytime there is saved changes in the workspace other than test files.

### Running the Worker

To run the worker, use the following command:
```bash
make run_worker
```

This command starts the boilerplate service worker.

To run the worker with live-reloading (during development only), use the following command:
```bash
make run_worker_dev
```
This command starts the boilerplate service worker that reload everytime there is saved changes in the workspace other than test files.

### Building the Project
To build the project, use the following command:

```bash
make build
```

This command compiles the Go code and generates the executable file in the `./bin/app` directory.

### Database Migrations

The boilerplate service supports database migrations. You can perform the following migration actions:

- Up: Apply all pending migrations.
```bash
make migrate_up
```

- Down: Rollback the most recent migration.
```bash
make migrate_down
```
- Up to a specific version: Apply migrations up to a specific version number.
```bash
make migrate_up version=20230707172200
```
- Down to a specific version: Rollback migrations down to a specific version number.
```bash
make migrate_down version=20230707172200
```
- make migrate_create name=update_users_table
```bash
make migrate_create name=update_users_table
```

### Changelog Generation

To generate a changelog for the project, use the following command:
```bash
make changelog
```
This command generates a changelog in the `CHANGELOG.md` file. By default, it includes all tags. You can specify a version to include only the changes since the specified version:
```bash
make changelog version=v1.0.0
```

### Linting

To run the linter on the project, use the following command:
```bash'
make lint
```

### Running Tests

To run the unit tests for the project, use the following command:
```bash
make test
```
This command also includes linting and provides the test coverage percentage.

To run only the unit tests for the project, use the following command:
```bash
make test_only
```
This command also provides the test coverage percentage if all test passed.

### Cleaning Mocks

To clean generated mock files, use the following command:

```bash
make clean_mock
```

This command removes all the `mock` directories from the project.

### Generating Mocks

To generate mock files, use the following command:

```bash
make mock
```

This command generates mocks based on the interfaces defined in the code.


### Generating Protobuf

To generate protobuf go files, use the following command:

```bash
make proto
```

This command generates golang protobuf based on the .proto defined in the code.
