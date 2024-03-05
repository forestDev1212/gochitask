## Start

#### Required\*

- GNU Make [Read Docs](https://www.gnu.org/software/make/)
- Docker [Read Docs](https://docs.docker.com/)
- Go version go1.21.5 [Read Docs](https://go.dev/dl/)

#### Setup project

- ##### Clone the Repository

  Make sure your git is setup using ssh-key and already in your code working directory

  ```
  git clone git@github.com:Kbgjtn/n.git
  cd n
  ```

- ##### Setup and Run server

  Make sure you're already have GNU make to run the command bellow:

  ```
  make setup
  make run
  ```

- ##### Test Server
  This will run all test code if you wanna check the tests
  ```
  make setup
  make run
  ```
- ##### Test API

  **Get Quote** (positive int)

  ```
  # GET /task/:id
  curl -i -X GET http://127.0.0.1:3000/api/task/1
  ```

  **Response**

  ```
  HTTP/1.1 200 OK
  Content-Type: application/json; charset=UTF-8
  Date: Fri, 29 Dec 2023 13:59:04 GMT
  Content-Length: 158

  {
    "code": 200,
    "message": "success",
    "data": {
        "id": 5,
        "title": "134",
        "priority": 123,
        "date": "2006-01-02T15:04:05Z",
        "created_at": "2024-03-04T14:56:10.187845Z",
        "updated_at": "2024-03-04T14:56:10.187845Z"
    }
}

  ```

  **List Quotes**

  ```
  # GET /tasks/:id (positive int)
  curl -i -X GET http://127.0.0.1:3000/api/tasks
  ```

  **Response**

  ```
  Content-Type: application/json; charset=UTF-8
  Date: Fri, 29 Dec 2023 14:06:35 GMT
  Content-Length: 569

  {
    "code": 200,
    "message": "success",
    "data": [
        {
            "id": 3,
            "title": "123",
            "priority": 123,
            "date": "2006-01-02T15:04:05Z",
            "created_at": "2024-03-04T14:50:49.11049Z",
            "updated_at": "2024-03-04T14:50:49.11049Z"
        },
        {
            "id": 4,
            "title": "123",
            "priority": 123,
            "date": "2006-01-02T15:04:05Z",
            "created_at": "2024-03-04T14:53:01.347691Z",
            "updated_at": "2024-03-04T14:53:01.347691Z"
        },
        {
            "id": 5,
            "title": "134",
            "priority": 123,
            "date": "2006-01-02T15:04:05Z",
            "created_at": "2024-03-04T14:56:10.187845Z",
            "updated_at": "2024-03-04T14:56:10.187845Z"
        },
        {
            "id": 6,
            "title": "123",
            "priority": 123,
            "date": "2006-01-02T15:04:05Z",
            "created_at": "2024-03-04T14:57:33.840414Z",
            "updated_at": "2024-03-04T14:57:33.840414Z"
        },
        {
            "id": 7,
            "title": "123",
            "priority": 123,
            "date": "2006-01-02T15:04:05Z",
            "created_at": "2024-03-04T14:58:37.653845Z",
            "updated_at": "2024-03-04T14:58:37.653845Z"
        }
    ],
    "length": 5,
    "paginate": {
        "offset": 0,
        "limit": 10,
        "total": 5,
        "prev": 0,
        "next": 5,
        "has_next": true,
        "has_prev": false
    }
}
  ```

  **Create Task**

  ```
  curl -v -XPOST -H "Content-type: application/json" \
  -d {
    "id" : 5,
    "title" : "123",
    "priority" : 134,
    "date" : "2006-01-02T15:04:05Z"
} \
  '127.0.0.1:3000/api/quotes'

  ```

  **Response**

  {
    "code": 200,
    "message": "success",
    "data": {
        "id": 8,
        "title": "134",
        "priority": 123,
        "date": "2006-01-02T15:04:05Z",
        "created_at": "2024-03-04T15:57:14.314585Z",
        "updated_at": "2024-03-04T15:57:14.314585Z"
    }
}

  ```

  **Task Quote** (id: positive int)

  ```
  curl -v -XPUT -H "Content-type: application/json" \
  -d '{
    "id" : 7,
    "title" : "123",
    "priority" : 134,
    "date" : "2006-01-02T15:04:05Z"
}'\  	'127.0.0.1:3000/api/quotes/4'
  ```

  **Response**

  ```
  Trying 127.0.0.1:3000...*
  Connected to 127.0.0.1 (127.0.0.1) port 3000
  > POST /api/quotes/4 HTTP/1.1
  > Host: 127.0.0.1:3000
  > User-Agent: curl/8.4.0
  > Accept: */*
  > Content-type: application/json
  > Content-Length: 69
  >
  < HTTP/1.1 200 OK
  < Content-Type: application/json; charset=UTF-8
  < Date: Fri, 29 Dec 2023 14:24:39 GMT
  < Content-Length: 151
  <* Connection #0 to host 127.0.0.1 left intact

 {
    "code": 200,
    "message": "success",
    "data": {
        "id": 7,
        "title": "134",
        "priority": 123,
        "date": "2006-01-02T15:04:05Z",
        "created_at": "2024-03-04T14:58:37.653845Z",
        "updated_at": "2024-03-04T14:58:37.653845Z"
    }
}
  ```

  **Delete Task**

  ```
  curl -v -XDELETE '127.0.0.1:3000/api/tasks/4'
  ```

  **Response**

  ```
  *   Trying 127.0.0.1:3000...
  * Connected to 127.0.0.1 (127.0.0.1) port 3000
  > DELETE /api/quotes/4 HTTP/1.1
  > Host: 127.0.0.1:3000
  > User-Agent: curl/8.4.0
  > Accept: */*
  >
  < HTTP/1.1 200 OK
  < Date: Fri, 29 Dec 2023 14:28:17 GMT
  < Content-Length: 0
  <
  * Connection #0 to host 127.0.0.1 left intact
  ```

## OpenAPI Doc

See Documentation REST API in here:
[Read Docs](https://github.com/Kbgjtn/n/tree/main/docs/swagger.json) or if you're running the server: [Go to Open API](http://localhost:3000/openapi)

## TODO

- [ ] Create an Open API Docs
- [ ] Cover all Test
- [ ] Create UI
