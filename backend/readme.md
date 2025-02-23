# Backend Setup Guide

This guide will help you set up the GoLang Gin and MongoDB backend locally.

## Prerequisites

- Go (version 1.23.1 or higher)
- MongoDB (version 4.4 or higher)
- Git

## Clone the Repository

```sh
git clone git@github.com:nayanrdeveloper/shopsocial.git
cd backend
```

## Setup .env File

Create a `.env` file in the root directory and add the following environment variables:

```env
MONGO_URI=mongodb://localhost:27017
DB_NAME=your_database_name
JWT_SECRET=any_key
PORT=8080
```

## Install Dependencies

Run the following command to install the required Go packages:

```sh
go mod tidy
```

## Run the Application

Use the following command to run the application locally:

```sh
go run cmd/main.go
```

The server should now be running at `http://localhost:8080`.

## Additional Commands

- To build the application:

    ```sh
    go build -o app
    ```

- To run tests:

    ```sh
    go test ./...
    ```

## Troubleshooting

- Ensure MongoDB is running locally.
- Verify the `.env` file contains the correct values.

For further assistance, refer to the official [GoLang Gin documentation](https://github.com/gin-gonic/gin) and [MongoDB documentation](https://docs.mongodb.com/).
