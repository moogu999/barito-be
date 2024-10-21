test:
	go test -v -timeout=60s -coverprofile=cover.out ./...

migration-create:
	migrate create -ext sql -dir migrations -seq $(NAME)
migration-migrate:
	migrate -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" -path migrations up
migration-rollback:
	migrate -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" -path migrations down $(STEPS)
migration-force:
	migrate -path migrations -database "mysql://$(USERNAME):$(PASSWORD)@tcp($(HOST):$(PORT))/$(DATABASE)" force $(VERSION)
