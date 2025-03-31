### API Methods

- Public
  - GET `/api/company/:id`, Get a signle Company by Name
  - POST `/api/auth/login`, Login using email and password, return a JWT token
  - POST `/api/auth/register`, Register using email and password
- JWT Protected
  - POST `/api/company/` - Create a new Company Entry 
  - PATCH `/api/company/:id` - Update an existing Company Entry (by Name) 
  - DELETE `/api/company/:id` - Delete a single Company by Name

### How to run with Docker
```
* Prerequisites
  - Docker and Docker Compose
```
#### *Automatically*
`sh run.sh`

#### *Or Manually*

In order to run the application with Docker, you have to:
1. Build the companies-microservice image: `docker build --tag companies-microservice .`
2. Run the compose file with the companies-microservice and the db (postgresql): `docker compose up -d`
3. Now the rest api is exposed to `localhost:8080` and can accept requests.

### Sample Requests to test the RestAPI

- The 'companies-microservice' is exposed to `localhost:8080`
- You can either use curl (more tedious) or Postman (recommended)
- For Postman: Import the postman file `docs/http_requests/Companies Microservices Requests.postman_collection.json`
  - An authorization bearer is required for "authenticated access".
  - Firstly register a user, the login to produce a jwt token, then copy this token in the "Authorization" section in postman by choosing "Bearer Token" in every requests that needs it (CreateCompany, DeleteCompany, UpdatedCompany).
- For curl: View `docs/curl_requests/curl.md`

#### For Postman
- Import the http_requests to Postman located in `./docs/https_requests`

#### For curl
- Use the requests located in `docs`

#### For overviewing the Database, DBBeaver was used. However, those are some useful queries to view the results from postgres container:
- The postgres user-password are located in `.env`. For dev purposes, it is `postgres`.
- View all entries:
  - In terminal that has docker access: `docker exec -it postgres-instance psql -U postgres -W challenge`
  - In psql: 
    - To view all tables: `\dt`
    - To quit: `\q`
    - For Users: `SELECT * FROM users;`
    - For Companies: `SELECT * FROM companies;`

### Tech Stack
- RestAPI : Fiber v2
- Database: Postresql
- Model Handling in Go: GORM
- Containerization: Docker 

### Implementation Plan (Roughly)

  - Create placeholders for files and functionalities - Done
  - Implement CRUD Functionality - Done 
  - Implement DB Driver - Done 
  - Implement "Auth" (JWT) - Done
  - Containerize - Done
  - Change Project's Layout - Done
  - Integration tests - TODO
  - Implement Kafka Driver - Not Done
