### Register User

- Body Structure
```
{
    "email": email,
    "password": string
}
```

- curl http request
```
curl --location 'localhost:8080/api/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@gmail.com",
    "password": "1234"
}'
```

### Login User

- Body Structure
```
{
    "email": email,
    "password": string
}
```

- curl http request
```
curl --location 'localhost:8080/api/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "test@gmail.com",
    "password": "1234"
}'
```

### Create a Company

- Body Structure
```
{
    "name": companyName,
    "description": companyDescription,
    "amount_of_employees": integer,
    "registered": bool,
    "company_type": "Corporations" or "NonProfit" or "Cooperative" or "Sole Proprietorship"
}
```

- Authorization header is required. Only a JWT from a logged in user is valid here.

- curl http request
```
curl --location 'localhost:8080/api/company' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM2NzUwODAsInVzZXJfaWQiOiI1YmIzYTUwNC1jMTdhLTRiNDgtYmVmMi1kNzAwMGE0NTFhMTYifQ.f2xl0V-xvZ45XXQ2qs-KTMYRv8sc9qUl3XOeyAiUCpQ' \
--data '{
    "name": "a_company",
    "description": "Company Description",
    "amount_of_employees": 10,
    "registered": true,
    "company_type": "Corporations"
}'
```

### Get Company by ID

- Parameter
`A uuid of a company is needed here`

- curl http request
```
curl --location 'localhost:8080/api/company/cbd7f6e0-b793-4957-929d-39f79986106c'
```


### Update Company by ID

- Body Structure (one or more fields, not all are required)
```
{
    "name": companyName,
    "description": companyDescription,
    "amount_of_employees": integer,
    "registered": bool,
    "company_type": "Corporations" or "NonProfit" or "Cooperative" or "Sole Proprietorship"
}
```

- Parameter
`A uuid of a company is needed here`

- Authorization header is required. Only a JWT from a logged in user is valid here.

- curl http request
```
curl --location --request PATCH 'localhost:8080/api/company/cbd7f6e0-b793-4957-929d-39f79986106c' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM2NzUwODAsInVzZXJfaWQiOiI1YmIzYTUwNC1jMTdhLTRiNDgtYmVmMi1kNzAwMGE0NTFhMTYifQ.f2xl0V-xvZ45XXQ2qs-KTMYRv8sc9qUl3XOeyAiUCpQ' \
--data '{
    "description": "Test???"
}'
```

### Delete Company by ID

- Parameter
`A uuid of a company is needed here`

- Authorization header is required. Only a JWT from a logged in user is valid here.

- curl http request
```
curl --location --request DELETE 'localhost:8080/api/company/60f9b776-f127-481b-9935-facea2fbe6e9' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDM2NzUwODAsInVzZXJfaWQiOiI1YmIzYTUwNC1jMTdhLTRiNDgtYmVmMi1kNzAwMGE0NTFhMTYifQ.f2xl0V-xvZ45XXQ2qs-KTMYRv8sc9qUl3XOeyAiUCpQ' \
--data ''
```