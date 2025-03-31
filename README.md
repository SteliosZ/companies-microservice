### RestAPI
- I'm more familiar with Fiber v2, so I will use this framework as the RESTAPI.
- Postgresql for database
- Gorm to handle the structs in Go
- Docker for containerization 

- Steps (Roughly):
  - Create placeholders for files and functionalities - Done
  - Implement CRUD Functionality - Done 
  - Implement DB Driver - Done 
  - Implement "Auth" (JWT) - Done
  - Containerize - Done
  - Change Project's Layout - TODO
  - Implement Kafka Driver - TODO
  - If time, integration tests - TODO

### API Methods

- Public
  - GET `/api/company/{companyName}`, Get a signle Company by Name
- JWT Protected
  - POST `/api/company/` - Create a new Company Entry 
  - PATCH `/api/company/{companyName}` - Update an existing Company Entry (by Name) 
  - DELETE `/api/company/{companyName}` - Delete a single Company by Name

### Docker Run
In order to run the application with docker, you have to:
1. Build the companies-microservice image: `docker build --tag companies-microservice .`
2. Run the compose file with both the companies-microservice and the db (postgresql): `docker compose up -d`
3. Now the rest api is exposed to localhost:8080 and can accept requests.
