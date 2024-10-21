run:
	go run ./cmd/
test:
	go test -v -timeout=60s -coverprofile=cover.out ./...

migration-create:
	migrate create -ext sql -dir migrations -seq $(NAME)
migration-migrate:
	migrate -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" -path migrations up
migration-rollback:
	migrate -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" -path migrations down $(STEPS)
# if an error occurs, the schema_version will be marked as dirty. once the issue is fixed, it needs to be forced so that other migrations can run again
migration-force:
	migrate -path migrations -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" force $(VERSION)
