.PHONY: clean test security build run

APP_NAME = gocommerce
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/platform/migrations
DATABASE_URL = postgresql://postgres:Mamahku46@database-1.cxr5z06jdt9q.ap-southeast-1.rds.amazonaws.com:5432/gocommerce?sslmode=disable
#DATABASE_URL = postgres://postgres:postgres@localhost:5432/go_fiber?sslmode=disable

go.run:
	go run ./main.go

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: swag build
	$(BUILD_DIR)/$(APP_NAME)

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

docker.run: migrate.up docker.compose.build docker.compose.up

docker.network:
	docker network inspect gofiber-network >/dev/null 2>&1 || \
	docker network create -d bridge gofiber-network

docker.network.connect:
	docker network connect gofiber-network gofiber-docker

docker.fiber.build:
	docker build -t gofiber-docker:1.0 .

docker.compose.build:
	docker-compose build

docker.compose.up:
	docker-compose up

docker.dev: docker.compose.build docker.network docker.compose.up

docker.fiber: docker.fiber.build
	docker run --name gofiber-docker \
		--network gofiber-network \
		-p 5000:5000 \
		gofiber-docker:1.0

docker.container:
	docker container stop gofiber-docker
	docker container rm gofiber-docker
	docker container create --name gofiber-docker -p 5000:5000 -e NAME=Docker gofiber-docker:1.0
#docker.postgres:
#	docker run --rm -d \
#		--name dev-postgres \
#		--network gofiber-network \
#		-e POSTGRES_USER=postgres \
#		-e POSTGRES_PASSWORD=password \
#		-e POSTGRES_DB=postgres \
#		-v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
#		-p 5432:5432 \
#		postgres

docker.stop: docker.stop.fiber

docker.stop.fiber:
	docker stop gofiber-docker

swag:
	swag init --parseDependency --parseInternal
