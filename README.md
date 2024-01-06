# Go REST API with SQLite

This is a Go REST API project with SQLite as the database.

## Features

- **Get All Users**: Retrieve a list of all users.
- **Get User by ID**: Retrieve a user by their ID.
- **Create User**: Create a new user.
- **Update User**: Update data of a user by their ID.
- **Delete User**: Delete a user by their ID.

## Technologies Used

- **Go**: Programming language for the backend.
- **SQLite**: Database for persistent storage.
- **Echo**: Web framework for building APIs in Go.

## Project Description

- `main.go`: Entry point for the application.
- `router/router.go`: Sets up all API routes with the Echo web framework.
- `api/user/user_handler.go`: Handles HTTP requests related to user operations.
- `api/user/user_route.go`: Defines API routes for user operations.
- `model/user.go`: Defines the User data model.
- `database/database.go`: Initializes and handles database connections.
- `service/user_service.go`: Implements business logic for user operations.
## Project Structure

```
.
├── main.go                 # Entry point for the application.
├── router
│   └── router.go           # Sets up API routes with the Echo web framework.
├── api
│   └── user
│       ├── user_handler.go # Handles HTTP requests related to user operations.
│       └── user_route.go   # Defines API routes for user operations.
├── model
│   └── user.go             # Defines the User data model.
├── database
│   └── database.go         # Initializes and handles database connections.
├── service
│   └── user_service.go     # Implements business logic for user operations.
└── config
    └── constants.go        # Defines project constants.
```
## Setup

1. Clone the repository.
2. Ensure you have Go installed on your machine.
3. Install dependencies
4. ```go run main.go ```




## Test 

* Please run the tests one by one.

Test Cases

1. Get All Users
Test method: TestGetAllUsers

This test fetches all users from the database and compares them with the expected user data.

2. Get User by ID
Test method: TestGetUserByID

This test fetches a user by their ID from the database and compares it with the expected user data.

3. Create User
Test method: TestCreateUser

This test creates a new user, fetches all users from the database, and verifies that the new user has been added.

4. Update User by ID
Test method: TestUpdateUserByID

This test updates a user by their ID, fetches all users from the database, and verifies that the user has been updated.

5. Delete User by ID
Test method: TestDeleteByID

This test deletes a user by their ID, fetches all users from the database, and verifies that the user has been deleted.
