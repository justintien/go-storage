# go-storage

A golang simple file/image service.

## Usage

```sh
docker-compose up -d --build
```

## Api

### test server alive

```sh
curl -X GET http://localhost:3001
# response: server alive!
```

### file upload

```http
POST /file/upload
```

#### 参数

Attribute | Type | Required | Description
--- | --- | --- | ---
`file` | file | yes | 档案

- file path will MD5(file), so 同一個檔案不會重複存二份

> example:

```sh
curl -X POST \
  http://localhost:3001/file/upload \
  -H 'content-type: multipart/form-data' \
  -F 'file=@file路徑'
```

##### response

```http
200 OK
```

```json
{
  "code": 0,
  "data": {
    "size": 190494,
    "path": "/file/45/75/4b/c0/f7/c3/38/bc/a5/53/0b/ed/0a/69/5d/a4.jpg"
  }
}
```

or failure

```http
400 Bad Request
```

```json
{
    "code": 400,
    "message": "上传错误"
}
```

### browser file

```http
GET /file/{path}
```

> example:

```text
browser: http://localhost:3001/file/45/75/4b/c0/f7/c3/38/bc/a5/53/0b/ed/0a/69/5d/a4.jpg
```

or

```http
404 Not Found
```

```text
404 page not found
```

### image service

> 若檔案為 .jpg,.jpeg,.png 將支緩圖片縮放(scale) & 品質(quality)調整

#### 参数

Attribute | Type | Required | Description
--- | --- | --- | ---
`q` | int | no | 品質比例(%, 100指的是不變, support jpg/jpeg only)
`s` | int | no | 縮放比例(%, 100指的是不變)

```http
GET /image/45/75/4b/c0/f7/c3/38/bc/a5/53/0b/ed/0a/69/5d/a4.jpg?q=50&s=50
```
