# URL Shortener

URL Shortener service written with Go and MongoDB. It provides the following functionalities.
* Short URL
* Respond to original URL when request short URL
* Access numbers of clicks on a short URL [To-DO]
* Multiple Short URL [TO-DO]

## Run

Run

``` make run ```

## API Endpoint
` GET localhost:8500/api/v1/redirect/{code}`

### Request
```json

```

### Response
```json
{
  "code": "Yzp3gtYGR",
  "url": "http://www.google.com",
  "click": 0,
  "created_at": "2021-01-30T18:25:35.797Z"
}
```

## API Endpoint

` POST localhost:8500/api/v1/redirect `

### Request

```json 
curl -X POST localhost:8500/api/v1/redirect 'Content-Type: application/json' -d '{"URL":"http://www.github.com"}'
```

### Response
```json
{
  "code": "Yzp3gtYGR",
  "url": "http://www.google.com",
  "click": 0,
  "created_at": "2021-01-30T21:25:35.797938525+03:00"
}
```