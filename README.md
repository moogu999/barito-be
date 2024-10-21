# `barito-be`

`barito-be` is a backend service for an online bookstore. Users can create accounts and start purchasing books.
 
# Overview

`barito-be` serves as the backend for an online bookstore. It handles user account creation, book searches, purchases, and order history management. The service uses MySQL 8 as its database.

# Development

You need a running MySQL database instance to run the service locally. The service uses [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) to convert OpenAPI 3.0 specifications into server-side implementations and HTTP models, aiming to reduce the need for writing boilerplate code.

You can simply run the `generate` command to generate the code:

```
go generate ./...
```

Don't forget to run `go mod tidy` in case any generated code requires a new package.

## `oapi-codegen`

You can define the OpenAPI 3.0 specifications in the `/api` directory, separated by domain. Once defined, you can place the generator in `/internal/{domain}/port/oapi`. Here's an example of a generator file:

```
package oapi

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --package $GOPACKAGE --generate=types,skip-prune -o types.gen.go ../../../../api/users.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --package $GOPACKAGE --generate=chi-server,strict-server -o server.gen.go ../../../../api/users.yaml
```

The generated code will be placed in the same directory as the generator. You can start using it from there.

## Migrations

The service uses [migrate](https://github.com/golang-migrate/migrate) to manage database migrations. All migrations are stored in the `/migrations` directory, making changes to the database easily trackable. The Makefile defines tasks for creating, migrating, and rolling back migrations

## Environment Variables

Sample values for the environment variables can be found in `sample.env`

<table>

<tr>
<th>
Name
</th>
<th>
Required
</th>
</tr>

<tr>
<td>
HTTP_PORT
</td>
<td>
No
</td>
</tr>

<tr>
<td>
SQL_USERNAME
</td>
<td>
Yes
</td>
</tr>

<tr>
<td>
SQL_PASSWORD
</td>
<td>
Yes
</td>
</tr>

<tr>
<td>
SQL_HOST
</td>
<td>
Yes
</td>
</tr>

<tr>
<td>
SQL_PORT
</td>
<td>
Yes
</td>
</tr>

<tr>
<td>
SQL_DATABASE_NAME
</td>
<td>
Yes
</td>
</tr>

<tr>
<td>
SQL_MAX_OPEN_CONS
</td>
<td>
No
</td>
</tr>

<tr>
<td>
SQL_CONN_MAX_LIFETIME
</td>
<td>
No
</td>
</tr>

<tr>
<td>
SQL_MAX_IDLE_CONS
</td>
<td>
No
</td>
</tr>

<tr>
<td>
SQL_CONN_MAX_IDLE_TIME
</td>
<td>
No
</td>
</tr>

</table>

## Test

For unit testing, you can run the `test` task in the `Makefile`. Integration testing is done using [venom](https://github.com/ovh/venom), and you can find the test specifications in the `/e2e/test` directory. To run the integration tests, you just need to execute this command from inside the `/e2e` directory:

```
docker compose -f .\e2e-compose.yaml up
```

The integration testing container has separate environment variables in case you want to run it alongside the regular service. After the tests are completed, you can also interact with the server via the integration testing container (it doesn't auto-stop).

## Run

To run the service, you can simply execute the `run` task in the `Makefile`.
