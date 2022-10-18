# Simple GO Lang REST API

> Simple RESTful API to create, read, update and delete. No database implementation yet

## Quick Start


``` bash
# Install mux router
go get -u github.com/gorilla/mux
```

``` bash
go build
./go-rest-api-tutorial
```

## Endpoints

### Get All Person
``` bash
GET /people
```
### Get Single Person
``` bash
GET /people/{id}
```

### Delete Person
``` bash
DELETE /people/{id}
```

### Create People
``` bash
POST /people

 Request sample
# {
#   "firstname":"John",
#   "lastname":"Doe",
#   "address":{"city":"London",  "state":"EN"}
# }
```

### Get University By Country Name
``` bash
GET /university
 Request sample
# Request Param add :
# country = turkey
```
