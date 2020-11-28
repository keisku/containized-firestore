# containized-firestore

Sample golang app with container firestore.

## How to start

```
$ make build
$ make up
$ make logs
```

## How to end

```
$ make down
```

## API Endpoints

### `GET /account`

Request

```
curl --location --request GET 'http://localhost:5000/account/'
```

Response

```
[{"user_id":"3c346b57-21bf-4514-b5df-857ef23ca030","account_id":"test001","mail":"test001@mail.com"},{"user_id":"42e677c7-9d46-4ef4-a9a4-040c38c27d36","account_id":"test002","mail":"test002@mail.com"},{"user_id":"830eed95-25b4-4342-9c3b-fbf8a054f607","account_id":"kskumgk63","mail":"keisuke.umegaki.630@gmail.com"},{"user_id":"fdc3edf5-bb49-4de1-abc6-1b36dd34a3e5","account_id":"test003","mail":"test003@mail.com"}]
```


### `GET /account/{id}`

Request

```
curl --location --request GET 'http://localhost:5000/account/830eed95-25b4-4342-9c3b-fbf8a054f607'
```

Response

```
{"user_id":"830eed95-25b4-4342-9c3b-fbf8a054f607","account_id":"kskumgk63","mail":"keisuke.umegaki.630@gmail.com"}
```


### `POST /account`

Request

```
curl --location --request POST 'http://localhost:5000/account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": "kskumgk63",
    "mail": "keisuke.umegaki.630@gmail.com"
}'
```

Response

```
{"user_id":"830eed95-25b4-4342-9c3b-fbf8a054f607","account_id":"kskumgk63","mail":"keisuke.umegaki.630@gmail.com"}
```


### `POST /account/{id}`

Request

```
curl --location --request POST 'http://localhost:5000/account/830eed95-25b4-4342-9c3b-fbf8a054f607' \
--header 'Content-Type: application/json' \
--data-raw '{
    "account_id": "update",
    "mail": "update630@gmail.com"
}'
```

Response

```
204 No Content
```

### `DELETE /account/{id}`

Request

```
curl --location --request DELETE 'http://localhost:5000/account/830eed95-25b4-4342-9c3b-fbf8a054f607'
```


Response

```
204 No Content
```
