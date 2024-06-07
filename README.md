# Blockchain Code Generator

## Getting Started with Dockerfile

To run the application using Dockerfile, follow these steps:

1. Build the Docker image:
    ```bash
    docker build -t bc-code-generator-app .
    ```

2. Run the container:
    ```bash
    docker run -p 8081:8080 --name bc-code-generator-app bc-code-generator-app
    ```

3. Run the container in the background:
    ```bash
    docker run -d -p 8081:8080 --name bc-code-generator-app bc-code-generator-app
    ```

4. Login to the container via Bash:
    ```bash
    docker exec -it bc-code-generator-app bash
    ```

5. Start the existing container:
    ```bash
    docker start bc-code-generator-app
    ```

6. Stop a container:
    ```bash
    docker stop bc-code-generator-app
    ```

7. Check running containers:
    ```bash
    docker ps
    ```

8. Inspect containers:
    ```bash
    docker ps -a
    ```

9. Remove the container:
    ```bash
    docker rm bc-code-generator-app
    ```

10. Delete an image:
    ```bash
    docker rmi bc-code-generator-app
    ```

## Run Test

To execute tests

1. Execute all test functions in your package and get detailed output, including test names, execution times, and results

    ```bash
    go test -v 
    ```

2. Execute all benchmark functions in your package and measure performance. It executes functions starting with 'Benchmark' and outputs results.

    ```bash
    go test -bench=.
    ```
