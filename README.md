# Golang Gin + SQLx Template

Full Golang API template with folder structure, database migrations and API routes. Also includes a optimized Dockerfile.

## Getting Started

For running this basic API you can use the follwing command and then access the API via http://localhost:8080

```bash
docker compose up --build
```

## MakeFile

Run build make command with tests

```bash
make all
```

Build the application

```bash
make build
```

Run the application

```bash
make run
```

## API Documentation

An API Response is ether success or failure. The following JSON Body will be returned

#### Success response

```json
{
    "message": "Custom Message",
    "data": any
}
```

#### Error response

```json
{
    "message": "Error Message"
}
```
