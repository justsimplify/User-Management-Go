# This is a sample user management project.

### Steps to Run:
- Clone the Repo.
- Create `config.yaml` in root directory as follows.
```yaml
host: 127.0.0.1
db:
  dbname: dbname
  username: dbuser
  password: dbpassword
```
- Make sure that the DB exists already in MySQL and MySQL runs on `3306` port.
- Run
`go run main.go`
This will run a webservice on `8888` port.

### API Docs
- GET /setup
  - Response:
    - error : If any error occurs during setup of table in DB.
    - response : 'Success' if everything executed successfully.

- POST /register - Add new user to db
  - Request Parameter:
    -
    ```json
    {"username": "user", "password": "password"}
    ```
  - Response:
    - error: If any error occurs during execution.
    - response: 'Success' if user was added successfully.

- POST /login - Login user
  - Request Parameter:
    -
    ```json
    {"username": "user", "password": "password"}
    ```
  - Response:
    - error: If any error occurs during execution.
    - response: JWT Token if user login was successful. This token is valid for 15 minutes only.

- POST /verifyToken - Verify token
  - Request Parameter:
    -
    ```json
    {"token": "JWT Token"}
    ```
  - Response:
    - error: If any error occurs during execution such as invalid token or expired token.
    - response: User Details if token is valid.
      ```json
      {
          "response": {
              "user_name": "username",
              "password": "password",
              "id": 1,
              "exp": 1577440103,
              "iss": "issuer_name"
          }
      }
      ```