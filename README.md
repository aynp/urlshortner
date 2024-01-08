# URL Shortner

A simple URL shortner built in Go.

## Usage

1. Run Redis
    ```bash
    docker run --name some-redis -p 6379:6379 -d redis redis-server --save 60 1 --loglevel warning
    ```

2. Run the server
    ```bash
    go run main.go
    ```

## APIs

### POST - `/v1/tiny/create`

Sample request body - 

```
{
  "original_url": "http://www.google.com",
  "url_params": [
    {
      "source_param": "x",
      "target_param": "y",
      "is_mandatory": true
    }
  ],
  "header_params": [
    {
      "source_param": "x",
      "target_param": "y",
      "is_mandatory": true
    }
  ],
  "auto_gen_params": [
    {
      "type": "uuid",
      "target_key": "uid"
    }
  ]
}
```

### GET - `/v1/tiny/redirect/{path}` 

Redirects to the original URL 

## TODO
1. Add tests
2. Document APIs better
