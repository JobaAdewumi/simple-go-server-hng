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
    "DeletedAt": null,
    
}

Example
POST url/api
Request
{
    "name": "Joba Adewumi"
}

Response 
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
```

