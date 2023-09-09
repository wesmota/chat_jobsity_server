DBNAME          = jobsity
TEST_DBNAME     = jobsity_test

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