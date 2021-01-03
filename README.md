# API REST with Go, Echo, GORM - API Movies

- EndPoints [/api/v1/authors]
    - /:ID [GET, DELETE]
    - / [POST, PUT]

```json
{
  "id": "number?",
  "name": "string",
  "last_name": "string", 
  "country": "string",
  "movies_id": "number"
}
```

- EndPoints [/api/v1/categories]
    - /:ID [DELETE]
    - / [GET, POST, PUT]

```json
{
  "id": "number?",
  "category": "string"
}
```

- EndPoints [/api/v1/movies]
    - / [POST, PUT]
    - /:ID [DELETE]
    - /all [GET]
    - /by/categories/:ID [GET]
    - /by/stars/:stars [GET]

```json
{
  "id": "number?",
  "name": "string",
  "stars": "number",
  "state": "number",
  "categories_id": "number"
}
```