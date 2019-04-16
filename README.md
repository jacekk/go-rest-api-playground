# go-rest-api-playground

### Requirements

- [Go](https://golang.org/doc/install) [ >=1.11 ]
- [docker-compose](https://docs.docker.com/compose/install/) [ >=3.* ] - optional; used for PostgreSQL as a DB 

### Running

- `cp dist.env .env` - and edit (if necessary)
- `docker-compose up -d` - optional; if PostgreSQL configured
- `go run cmd/server.go`
