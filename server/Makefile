postgresinit: 
	docker run --name postgres16 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

postgres:
	docker exec -it postgres16 psql

createdb: 
	docker exec -it postgres16 createdb --username=root --owner=root chatigo

dropdb: 
	docker exec -it postgres16 dropdb chatigo

migrateup:
	migrate -path db/migrations -database "postgresql://root@root@localhost:5433/chatigo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root@root@localhost:5433/chatigo?sslmode=disable" -verbose down


.PHONY: postgresinit postgres createdb dropdb migrateup migratedown
