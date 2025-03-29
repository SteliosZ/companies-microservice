#### Keeping log before finalizing...

### RestAPI
- I'm more familiar with Fiber v2, so I will use this framework as the RESTAPI.
  
`go get github.com/gofiber/fiber/v2`

- Microservice Should have:
  1. Create
  2. Patch
  3. Delete
  4. Get (one)

- Company Model
  - ID (uuid) *required*
  - Name (15 characters) *required - unique*
  - Amount of Employees (int) *required*
  - Registered (boolean) *required*
  - Type (Corporations | NonProfit | Cooperative | Sole Proprietorship) *required*
  - Description (3000 characters) *optional*

- Considered Plus:
  - On each mutating operation, an event should be produced
  - Dockerize the application to be ready for building the production docker image
  - Use docker for setting up the external services such as the database
  - REST is suggested, but GRPC is also an option
  - JWT for authentication
  - Kafka for events
  - DB is up to you (will postgresql)
  - Integration tests are highly appreciated
  - Linter
  - Configuration file
