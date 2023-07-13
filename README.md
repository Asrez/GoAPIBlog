# Blogging System API

This project is a RESTful API built in Go language that serves as the backend for a blogging system. It provides endpoints for creating, retrieving, updating, and deleting blog posts.

## Features

- User authentication and authorization using JSON Web Tokens (JWT)
- CRUD operations for blog posts
- Pagination and sorting of blog posts
- Error handling and response formatting
- Input validation and data sanitization
- Logging and request/response middleware
- Database integration using PostgreSQL

## Installation

Make sure you have Go installed on your system. You can download and install it from the official Go website: https://golang.org

Clone the repository:

```bash
git clone https://github.com/BaseMax/GoAPIBlog.git
```

Change to the project directory:

```bash
cd GoAPIBlog
```

Install the project dependencies:

```bash
go mod download
```

Set up the environment variables:

- Create a `.env` file in the root directory of the project.

- Define the following environment variables in the `.env` file:

```plaintext
PORT=3000
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
JWT_SECRET=your_jwt_secret
```

Build and run the application:

```bash
go build
./blogging-api
```

The API will be accessible at http://localhost:3000 by default. You can change the port in the .env file if necessary.

## API Endpoints

The following endpoints are available in the API:

| Method | 	Endpoint | 	Description |
| ---- | -------- | -------- |
| POST |	/api/auth/register	| Register a new user |
| POST |	/api/auth/login	| Log in and obtain JWT |
| GET |	/api/posts	| Get all blog posts |
| POST |	/api/posts	| Create a new blog post |
| GET |	/api/posts/:id	| Get a specific blog post |
| PUT |	/api/posts/:id	| Update a specific blog post |
| DELETE |	/api/posts/:id	| Delete a specific blog post |

## Database Schema

The application uses a PostgreSQL database with the following schema:

```sql
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    author_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);
```

## Dependencies

The project utilizes the following third-party libraries:

- `gin-gonic/gin`: HTTP web framework
- `joho/godotenv`: Environment variable loading
- `dgrijalva/jwt-go`: JWT implementation
- `go-pg/pg`: PostgreSQL ORM
- `go-pg/pgext`: PostgreSQL extensions for go-pg
- `google/uuid`: UUID generation
- `go-playground/validator/v10`: Struct field validation

Make sure to run go mod download as mentioned in the installation steps to fetch these dependencies.

## Contributing

Contributions to this project are welcome. Feel free to open issues and submit pull requests.

## License

This project is licensed under the GPL-3.0 License.

# Postman collection

[postman](https://documenter.getpostman.com/view/20769678/2s946e9Ynx)

Copyright 2023, Max Base

Copyright 2023, Asrez group
