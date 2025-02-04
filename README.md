# Mail Verifier

This is a Go project that can be Dockerized, allowing it to run in a containerized environment. This README explains how to set up, build, and run the project with Docker. It also includes instructions for passing environment variables such as `PORT` and `PROXY_URI` during the build process.

## Prerequisites

Before getting started, ensure you have the following installed:

- [Docker](https://www.docker.com/get-started) (for building and running containers)
- [Go](https://golang.org/doc/install) (for building the project locally if needed)

## Dockerizing the Go Project

### Dockerfile Overview

This project includes a `Dockerfile` that allows you to build a Docker image for the Go application. Here's a breakdown of the `Dockerfile`:

1. **Base Image**: We use the official `golang:1.20-alpine` image as the base image. You can change the Go version if needed.
2. **Environment Variables**: The `PORT` and `PROXY_URI` environment variables are set inside the Dockerfile. The `PORT` is set to `3000` by default.
3. **Build and Run**: The Dockerfile performs the following steps:
   - Sets up the Go environment.
   - Copies the Go modules and source code into the container.
   - Downloads the Go modules.
   - Builds the Go application.
   - Exposes the port specified in the environment variable `PORT`.

## Installation
building the docker image

```
docker build --build-arg PROXY_URI=<YOUR PROXY> -t go-app .
```
run the image

```
docker run -p 3000:3000 go-app
```


## API Documentation

### Overview

The Go application exposes several endpoints for email verification. It uses the Gin framework to handle HTTP requests. Below is the detailed API documentation for each endpoint.

### Base URL
The base URL for the API is: http://localhost:3000


### Routes

#### 1. `GET /mail`
This is a simple test route to check if the server is running.

**Response:**
- **Status:** `200 OK`
- **Body:**
```json
{
  "message": "Success",
  "data": "yo!!",
  "success": true
}
````

#### 2. `POST /mail`
This is the verify a single mail.

**Body**
```json
{
    "email":"anshuman9998@gmail.com",
    "use_proxy":true {Optional}
}

```

**Response:**
- **Status:** `200 OK`
- **Body:**
```json
{
    "data": {
        "email": "anshuman9998@gmail.com",
        "reachable": "unknown",
        "syntax": {
            "username": "anshuman9998",
            "domain": "gmail.com",
            "valid": true
        },
        "smtp": null,
        "gravatar": null,
        "suggestion": "",
        "disposable": false,
        "role_account": false,
        "free": true,
        "has_mx_records": true
    },
    "error": null,
    "success": false
}
```

#### 3. `POST /mail/bulk`
This is to verify bulk mails.

**Body**
```json
{
    "emails": ["anshuman9998@gmail.com","bcabncanshuman2020@gmail.com"],
    "use_proxy":true {Optional}
}

```

**Response:**
- **Status:** `200 OK`
- **Body:**
```json
{
   {
    "data": {
        "anshuman9998@gmail.com": {
            "email": "anshuman9998@gmail.com",
            "reachable": "unknown",
            "syntax": {
                "username": "anshuman9998",
                "domain": "gmail.com",
                "valid": true
            },
            "smtp": null,
            "gravatar": null,
            "suggestion": "",
            "disposable": false,
            "role_account": false,
            "free": true,
            "has_mx_records": true
        },
        "bcabncanshuman2020@gmail.com": {
            "email": "bcabncanshuman2020@gmail.com",
            "reachable": "unknown",
            "syntax": {
                "username": "bcabncanshuman2020",
                "domain": "gmail.com",
                "valid": true
            },
            "smtp": null,
            "gravatar": null,
            "suggestion": "",
            "disposable": false,
            "role_account": false,
            "free": true,
            "has_mx_records": true
        }
    },
    "error": null,
    "success": false
}
```




