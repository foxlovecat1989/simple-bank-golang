### docker ###
dockerup:
	docker compose up -d
dockerdown:
	docker compose down

### DB ###
# docker exec -it ${CONTAINER_NAME_OR_ID} createdb --username=${USERNAME} --owner=${OWNER} ${DB_NAME}
createdb:
	docker exec -it postgres createdb --username=ed --owner=ed simple_bank
# docker exec -it ${CONTAINER_NAME_OR_ID} dropdb --username=${USERNAME} ${DB_NAME}
dropdb:
	docker exec -it postgres dropdb --username=ed simple_bank

### migrate ###
migrateup:
	migrate -path db/migration -database "postgresql://ed:password@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://ed:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

### sqlc ###
sqlc:
	sqlc generate

.PHONY: dockerup dockerdown createdb dropdb migrateup migratedown sqlc