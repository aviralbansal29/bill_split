# Bill Split
Clone of Splitwise

## Creating Environment:
- Update env as needed in env.yaml
- Environment variables can be set to overwrite env.yaml, for eg. `postgres.db` will be overwritten by environment variable `POSTGRES_DB`

## Server:
- Build server  : `docker-compose build`
- Run server    : `docker-compose up -d`
- Stop server   : `docker-compose down`

## Database config:
- Update postgres env varaibles inside .env.docker
- Varaibles to Set : 
  - POSTGRES_USER=<DB User>
  - POSTGRES_PORT=<DB Port>
  - POSTGRES_PASSWORD=<DB password>
  - POSTGRES_DB=<DB name>
  - POSTGRES_HOST=database

## Swagger
After strting server, docs can be found here http://localhost:9000/swagger/index.html#/
