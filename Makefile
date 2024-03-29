DBNAME          = jobsity
TEST_DBNAME     = jobsity_test
LOCAL_DB_ENV    = DB_NAME=${DBNAME} DB_PORT=$${DB_PORT:-5432} DB_HOST=localhost DB_USER=$${PGUSER:-root} DB_PASSWORD=$${PGPASSWORD:-password}
GO_BUILD        = env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"
JWT_TTL		 	= 60
JWT_SECRET	 	= secret
JWT_ISSUER	 	= jobsity
RMQ_HOST = localhost
RMQ_USERNAME = guest
RMQ_PASSWORD = guest
RMQ_PORT = 5672
LOCAL_RMQ_ENV = RMQ_HOST=${RMQ_HOST} RMQ_USERNAME=${RMQ_USERNAME} RMQ_PASSWORD=${RMQ_PASSWORD} RMQ_PORT=${RMQ_PORT}
LOCAL_JWT = JWT_TTL=${JWT_TTL} JWT_SECRET=${JWT_SECRET} JWT_ISSUER=${JWT_ISSUER}

db-create:
	PGPASSWORD=password psql -h localhost -U root -f db/setup/local/create.sql
	
db-run:
	docker-compose up -d
	until pg_isready -h localhost; do sleep 1; done

db-build-from-schema:
	PGPASSWORD=password psql -h localhost -U root -f db/schema.sql ${DBNAME} > /dev/null

db-seed:
	PGPASSWORD=password psql -h localhost -U root -f db/setup/local/seed.sql ${DBNAME} > /dev/null

db-reset: db-run db-create db-migrate db-seed

db-reset-no-seed: db-run db-create db-build-from-schema

goose-status: 
	goose postgres "user=root password=password dbname=${DBNAME} sslmode=disable" status

db-migrate:
	goose postgres "user=root password=password dbname=${DBNAME} sslmode=disable" up
	pg_dump postgres://root:password@localhost:5432/${DBNAME} --schema-only --no-owner --file db/schema.sql
run: db-reset
	export ${LOCAL_RMQ_ENV} && \
	export ${LOCAL_JWT} && \
	go mod tidy && \
	go mod download && \
	go run main.go

tests:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

