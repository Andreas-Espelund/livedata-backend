# Livedata Backend

This project provides a RESTful API implemented in Go. It utilizes the Swagger/OpenAPI 2.0 specification for API documentation.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed the latest version of [Go](https://golang.org/dl/) and [Docker](https://www.docker.com/products/docker-desktop).
- You have installed [Go Swagger](https://github.com/go-swagger/go-swagger). To install, run:
    ```bash
    go get -u github.com/go-swagger/go-swagger/cmd/swagger
    ```

## Setting Up

To set up the Livedata Backend, follow these steps:

1. Clone this repository:
    ```bash
    git clone https://github.com/your_username/livedata-backend
    ```

2. Navigate to the project directory:
    ```bash
    cd livedata-backend
    ```

3. Generate the server from the Swagger specification:
    ```bash
    make generate
    ```

## Running the Server

To run the server, execute the following command in your terminal:

```bash
make server
```

The server will start at `localhost:8080`.

## Docker

You can also build and run this project as a Docker container:

1. Build the Docker image:
    ```bash
    docker build -t livedata-backend .
    ```

2. Run the Docker container:
    ```bash
    docker run -p 8080:8080 livedata-backend
    ```

The server will be available at `localhost:8080`.

## API Documentation

You can view the API documentation at `localhost:8080/swagger`.

## Contributing

To contribute to Livedata Backend, follow these steps:

1. Fork this repository.
2. Create a branch: `git checkout -b <branch_name>`.
3. Make your changes and commit them: `git commit -m '<commit_message>'`.
4. Push to the original branch: `git push origin <project_name>/<location>`.
5. Create the pull request.

## Contact

If you want to contact me, you can reach me at `<your_email>`.

## License

This project uses the following license: `<license>`.

---

Note: Replace `<your_username>`, `<your_email>`, `<license>`, `<branch_name>`, `<commit_message>`, and `<project_name>/<location>` with appropriate values.