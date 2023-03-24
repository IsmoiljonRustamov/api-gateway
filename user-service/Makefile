CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

gen-proto:
	./scripts/gen-proto.sh

tidy:
	go mod tidy
	go mod vendor
run:
	go run cmd/main.go	

create:
	migrate create -ext sql -dir ./migrations -seq create_user_table   

up:
	migrate -path migrations/ -database postgres://ismoiljon12:12@localhost:5432/user_db up


down:
	migrate -path migrations/ -database postgres://ismoiljon12:12@localhost:5432/user_db down


force:
	migrate -path migrations/ -database postgres://ismoiljon12:12@localhost:5432/user_db force 1