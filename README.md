# My first go API

## To run this APP: 

- Clone: `git clone https://github.com/SouzaBernardo/first-go-api.git`
- Go to the folder that cloned the repo: `cd first-go-api/`

### With Docker:

- Docker-compose up
- Access [localhost:8080](http://localhost:8080/)
  
#### PGAdmin:
- You can access `localhost:16543/` and log in with `PGADMIN_DEFAULT_EMAIL` and `PGADMIN_DEFAULT_PASSWORD` to manipulate PostgreSQL.
- To register a server:
	- Rigth click on `Servers` > `Register` > `Servers`
	- On section `Connection` set host name to `postgres-container`
	- If you don't change this values:
		- Username: `postgres`
		- password: `postgres`

### Without Docker:

- Install go 1.21
- Go to the folder that cloned the repo: `cd first-go-api/`
- Run `go mod tidy`
- Install [PostgreSQL](https://www.postgresql.org/download/)
- Open PGAdmin and run the schema on `first-go-api/model/scripts/schema.sql`
- Run `go run main.go` in `/first-go-api`

Technical inforations:
- go 1.21 version
- [pq package](https://pkg.go.dev/github.com/lib/pq)
- docker and docker-compose
