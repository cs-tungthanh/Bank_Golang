# Tools Instruction

## 1. [Install golang-migrate](https://github.com/golang-migrate/migrate)
```
brew install golang-migrate
```
We're going to use 3 commands:
- `create`: we can use to create a new migration file.
- `goto`: which will migrate the schema to a specific version
- `up` or `down`: to apply all or N up or down migrations.

Example:
```
migrate create -ext sql -dir db/migration -seq init_schema
```

## 2. [Install sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)
```
brew install sqlc
```

## 3. Install mockgen
```
go install github.com/golang/mock/mockgen@v1.6.0
export PATH=$PATH:~/go/bin
```

Why mock database?
- independent test: avoid conflicts with real data
- faster test:
    we don't need to talk to the database
    all actions will be performed in memory, on the same process
- 100% coverage
    easily setup edge cases: unexpected errors or connection lost
     
How: using stubs instead of fake db
    - fake db(implement a fake version of DB)
-> Using Gomock to mock db

```
mockgen pathToModule/db/sqlc InterfaceName
Example: mockgen -package mockdb -destination db/mock/store.go github.com/cs-tungthanh/Bank_Golang/db/sqlc store
```

## Docker
`docker container inspect postgres12`: to see the network setting
- check NetworkSetting/IPAddress

```bash
docker network ls
docker network create network-name

# to list all containers are running in this network
docker network inspect network-name

# Manual connect: 
docker network connect network-name container
```

- If you have more than 2 containers with non-defined network, it will run in different networks.
- That means each container will run with different IPAddress and cannot connect to each other.
- so we are supposed to attach this container to the same network.
- if all services are defined in the same file docker-compose, it has the same network
- the network is just need when we run seperately container
networks
### Tech
- need to create network first: docker network create bank-network
```
networks:
  bank-network:
    external: true
```

```bash
docker exec -it container-name bash
psql -U root -d simple_bank
```

## References
- https://github.com/golang-migrate/migrate#cli-usage