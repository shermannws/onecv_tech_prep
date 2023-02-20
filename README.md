# Golang Assessment for OneClientView Tech Role
Done by Sherman Ng

## Project Structure
- `config/` contains the MySQL database configuration code
- `controller/` contains the business logic for the respective API endpoints
- `db-setup/` contains the SQL script for setting up the database
- `model/` contains the ORM model and database logic
- `postman/` contains the postman collection used for personal testing
- `test/` contains the unit test script for the application
- `main.go` is the entry point for the application

## Running the application locally
1. Download Go 1.20.1 from https://go.dev/dl/
2. Download MySQL 8.0.32.0 from https://dev.mysql.com/downloads/installer/
3. Set up MySQL using the script provided in `db-setup/setup.sql`
4. Install dependencies in project root directory
> `go get -u github.com/go-sql-driver/mysql` \
> `go get github.com/gorilla/mux`
5. Configure your database information in `./config/config.go` for `dbUser` and `dbPass`
6. Start the programme with command `go run ./main.go`
7. Local host is started on `http://localhost:8000`

## Running Unit Test
1. In a terminal, change directory into the test folder using `cd ./test`
2. Run `go test -v` to run all test cases
