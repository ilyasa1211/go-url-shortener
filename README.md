# URL Shortener in Go

This is a basic example of URL Shortening Application in Go, using Sqlite3 as the database.

## Run Locally

You have to install these on your machine, or if you had docker installed look section below.

- Go
- Sqlite3

## Run inside Docker

Build the image

```bash
docker build --tag url-shortener .
```

Run

```bash
docker run -p 8080:8080 -d url-shortener
```

## Prequisities

You need to create the database first using sqlite3

Create database

```bash
sqlite3 data/data.db
```

Create table `sites`

```bash
CREATE TABLE sites (id INTEGER PRIMARY KEY AUTOINCREMENT, alias_url VARCHAR(255) NOT NULL, target_url VARCHAR(255) NOT NULL);
```

Now you're good to go

## API Reference

#### Get all stored urls

```http
  GET http://localhost:8080/sites
```

#### Get a target url

You will get auto redirection if you enter this url

```http
  GET http://localhost:8080/sites/{aliasUrl}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `aliasUrl`      | `string` | **Required**. the short url version |

#### Create a short url

```http
  POST http://localhost:8080/sites
  Content-Type: application/json
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `alias_url`      | `string` | **Required**. the short url version |
| `target_url`      | `string` | **Required**. the target url |

#### Update a short url

This one is for update the target url

```http
  PUT http://localhost:8080/sites/{aliasUrl}
  Content-Type: application/json
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `target_url`      | `string` | **Required**. the new target url |

#### Delete a short url

```http
  DELETE http://localhost:8080/sites/{aliasUrl}
```
