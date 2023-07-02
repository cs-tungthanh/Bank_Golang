createNetwork:
	docker network create bank-network

run:
	docker-compose up -d

postgres:
	docker run -d \
		-e POSTGRES_USER=root \
		-e POSTGRES_PASSWORD=secret \
		-e POSTGRES_DB=simple_bank \
		-p 5432:5432 \
		--name postgres12 \
		postgres:12-alpine 

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

# modify localhost to db'hostname and password to access to db
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

# run all unittests in all packages
test:
	go test -v -cover ./...

server:
	go run main.go

dockerbuild:
	docker build -t simplebank:latest .


#  because we have postgres12 in the same network so hostname can use as the container name in DB_SOURCE
dockerrun:
	docker run --name simplebank \
		-p 8080:8080 \
		--network bank-network \
		-e GIN_MODE=release \
		-e DB_SOURCE="postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable" \
		simplebank:latest

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/cs-tungthanh/Bank_Golang/db/sqlc Store      

.PHONY: postgres createdb dropdb migratedown migrateup sqlc test server mock
