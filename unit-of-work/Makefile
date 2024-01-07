createmigration:
	migrate create -ext=sql -dir=sql/migrations -seq $(name)

migrate: 
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose up
	
migratedown: 
	migrate -path sql/migrations -database "mysql://root:root@tcp(localhost:3306)/courses" -verbose down

generatesql:
	sqlc generate

.PHONY: migrate createmigrataion migratedown generatesql
