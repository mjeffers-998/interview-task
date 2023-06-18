### Implementation

* SQLite database backend
* gin as the web framework

To run the server, do 

```
go run cmd/main.go
```

To make a POST request, use a multipart form - as an example:

```
curl --request POST \
  --url http://localhost:9007/orders \
  --header 'Content-Type: multipart/form-data' \
  --form order_date=2024-12-31T21:07:14-05:00 \
  --form order_value=3243 \
  --form store_id=423 \
  --form status=new
```


### COMPLETED:
* Database Schema
* Storage interface fully implemented for SQLite
* API implemented for CreateOrder

### TODO
* API implementation for all other calls (just stubs currently)
* Unit testing/load testing
* Validation middleware