createdb:
	docker exec -it postgres12_wallet createdb --username=root --owner=root simple_wallet

createtestdb:
	docker exec -it postgres12_wallet createdb --username=root --owner=root simple_wallet_test

dropdb:
	docker exec -it postgres12_wallet dropdb simple_wallet

droptestdb:
	docker exec -it postgres12_wallet dropdb simple_wallet_test

migrateup:
	docker exec -it api_wallet /app/migrate -path /app/migration -database "postgresql://root:secret@wallet.database.loc:5432/simple_wallet?sslmode=disable" -verbose up

migratetestup:
	docker exec -it api_wallet /app/migrate -path /app/migration -database "postgresql://root:secret@wallet.database.loc:5432/simple_wallet_test?sslmode=disable" -verbose up

migrateup1:
	docker exec -it api_wallet /app/migrate -path /app/migration -database "postgresql://root:secret@wallet.database.loc:5432/simple_wallet?sslmode=disable" -verbose up 1

migratedown:
	docker exec -it api_wallet /app/migrate -path /app/migration -database "postgresql://root:secret@wallet.database.loc:5432/simple_wallet?sslmode=disable" -verbose down

migratetestdown:
	docker exec -it api_wallet /app/migrate -path /app/migration -database "postgresql://root:secret@wallet.database.loc:5432/simple_wallet_test?sslmode=disable" -verbose down

migratedown1:
	docker exec -it api_wallet /app/migrate -path /app/migration -database "postgresql://root:secret@wallet.database.loc:5432/simple_wallet?sslmode=disable" -verbose down 1

sqlc:
	docker run --rm -v ${PWD}:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/MaksimDzhangirov/wallet/db/sqlc Store

.PHONY: createdb dropdb createtestdb droptestdb migrateup migratetestup migratedown migratetestdown migrateup1 migratedown1 sqlc test server mock