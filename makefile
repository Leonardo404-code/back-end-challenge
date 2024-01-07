#!bash

db_up:
	goose -dir ./migrations postgres "host=localhost sslmode=disable password=postgres dbname=packages user=postgres port=9948" up