ENV_FILE=.env

include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

.PHONY: setup-db
setup-db:
	cd heroes_cube; go mod tidy;
	cd heroes_cube/scripts; go run migrate.go
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/creatures.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/races.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/classes.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/inventories.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/items.sql;

.PHONY: setup
setup: setup-db
	@echo "Database setup complete."

.PHONY: test
test:
	cd heroes_cube/models; go clean -testcache; go test . -v;

.PHONY: build
build:
	cd heroes_cube; go mod tidy; go build -o heroes_cube main.go; chmod +x heroes_cube;

.PHONY: run
run:
	cd heroes_cube; ./heroes_cube;