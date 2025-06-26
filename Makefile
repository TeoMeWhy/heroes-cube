ENV_FILE=.env

include $(ENV_FILE)
export $(shell sed 's/=.*//' $(ENV_FILE))

.PHONY: setup-db
setup-db:
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/people.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/races.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/classes.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/inventories.sql;
	mysql -u $$DB_USER -p$$DB_PASS -h $$DB_HOST -D $$DB_NAME < data/items.sql;

.PHONY: setup
setup: setup-db
	@echo "Database setup complete."