# Blockchain Code Generator

## Getting Started with Dockerfile

To run the application using Dockerfile, follow these steps:

1. Build the Docker image:
    ```bash
    docker build -t bc-code-generator-app .
    ```

2. Run the container:
    ```bash
    docker run -d --name bc-code-generator-app bc-code-generator-app
    ```

3. Login to the container via Bash:
    ```bash
    docker exec -it bc-code-generator-app bash
    ```

## Getting Started with docker-compose

To run the application using docker-compose, follow these steps:

1. Run the app in detached mode:
    ```bash
    docker-compose up -d
    ```

## Build App

To build the application from the source code, follow these steps:

1. To build the application, run the following command:
    ```bash
    go build -o ./build/bc-code-generator-app
    ```

## Run App

To run the application, you can choose one of the following methods:

1. Run the application using the built binary:
    ```bash
    ./build/bc-code-generator-app
    ```

2. Run the application without building the binary file:
    ```bash
    go run main.go
    ```

## Run Test

To execute tests

1. Execute all test functions in your package and get detailed output, including test names, execution times, and results
    ```bash
    go test -v
    ```

2. Execute all benchmark functions in your package and measure performance. It executes functions starting with 'Benchmark' and outputs results
    ```bash
    go test -bench=.
    ```
