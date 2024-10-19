package oapi

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --package $GOPACKAGE --generate=types,skip-prune -o types.gen.go ../../../../api/users.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --package $GOPACKAGE --generate=chi-server,strict-server -o server.gen.go ../../../../api/users.yaml
