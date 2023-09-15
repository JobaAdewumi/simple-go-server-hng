## Documentation file for the server

### Standard formats for endpoints
Create
```bash
api/

Request 
{
    "name": your name
}

Response 
{
    "id": Your id
    "name": "Your name"
    "createdAt": "time"
    "UpdatedAt": "time"
    "DeletedAt": "time" || "null",
    
}

Example
POST url/api
Request
{
    "name": "Joba Adewumi"
}

Response 201
{
    "id": 1
    "name": "Joba Adewumi"
    "CreatedAt": "2023-09-14T20:19:01.814394+01:00",
    "UpdatedAt": "2023-09-14T20:19:01.814394+01:00",
    "DeletedAt": null,
}

```

Get
```bash
api/user_id

Request 
user_id as url parameter

Response 
{
    "id": Your id
    "name": "Your name"
    "createdAt": "time"
    "UpdatedAt": "time"
}

Example
GET url/api/:user_id
Request
user_id from url parameter

Response 200
{
    "id": 1
    "name": "Joba Adewumi"
    "CreatedAt": "2023-09-14T20:19:01.814394+01:00",
    "UpdatedAt": "2023-09-14T20:19:01.814394+01:00",
    "DeletedAt": null,
}
```

Update
```bash
api/user_id

Request 
{
    "name": your name
}

Response 
{
    "Rows Affected": rows
    
}

Example
PUT url/api/:user_id
Request
user_id from url parameter and json body
{
    "name": "Joba Adewumi"
}

Response 200
{
    "Rows Affected": 1
    
}
```
Delete
```bash
api/user_id

Request 
user_id from url parameter

Response 
{
    "success": "Delete successful"
    
}

Example
DELETE url/api/:user_id
user_id from url parameter

Response 200
{
    "success": "Delete successful"
    
}
```
## Run Server
### Prerequisites
```bash
go version >= 1.20.5
postgresql >= 14
```