CWD=$$(pwd)

build_image:
	docker build -t amino_acid_calc .

dev:
	docker run --network host --rm -it -v "${CWD}":/amino-acid-calc/ -w /amino-acid-calc amino_acid_calc

build:
	go build -o amino-acid-calc-core

test:
	go test

# DB MIGRATIONS
create_migration:
    # make create_migration MIG_NAME=some_mig_name
	docker run -v "${CWD}"/db/migrations/:/migrations --network host migrate/migrate create -ext sql -dir /migrations -seq "${MIG_NAME}"
	sudo chmod 777 -R "${CWD}"/db/migrations/

run_migration_up:
    # DATABASE URL FORMAT: dbdriver://username:password@host:port/dbname?option1=true&option2=false
	docker run -v "${CWD}"/db/migrations/:/migrations --network host migrate/migrate -database postgres://postgres:mypassword@localhost:5432/amino_acid_db?sslmode=disable -path /migrations -verbose up

run_migration_down:
    # DATABASE URL FORMAT: dbdriver://username:password@host:port/dbname?option1=true&option2=false
	docker run -v "${CWD}"/db/migrations/:/migrations --network host migrate/migrate -database postgres://postgres:mypassword@localhost:5432/amino_acid_db?sslmode=disable -path /migrations -verbose down -all
