# Link Shortener Service

Very simple link shortener service written in Go (Golang).

## How to run

1. Create `.env` file and

```shell
cp .env.example .env
```

2. Run container

```shell
docker-compose up -d
```

## Endpoints

### Ping server

```
GET /ping
```

#### Response

HTTP 200 OK
```json
{
  "message": "pong"
}
```

### Create short URL

```
POST /links
```

#### Parameters

| Name       | Required | Type    | Description                                                                                       |
|------------|----------|---------|---------------------------------------------------------------------------------------------------|
| url        | Yes      | string  | URL to shorten.<br>Must be a valid UR.L                                                           |
| single_use | No       | integer | Whether the short url can be used only once or not.<br>0: multiple use (default)<br>1: single use |


#### Response

HTTP 200 OK
```json
{
  "key": "1e1df7dd",
  "url": "http://google.com"
}
```

### Redirect to URL

```
GET /:key
```

#### Response

HTTP 301 Moved Permanently

## License

This project is open-sourced software licensed under the [MIT license](https://opensource.org/licenses/MIT).
