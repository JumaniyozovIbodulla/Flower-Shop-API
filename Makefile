run:
	go run main.go

push:
	git push origin main

swag_init:
	swag init -g api/router.go -o api/docs

migfile:
	migrate create -ext sql -dir migrations/ -seq create_tables

# path:
# 	export PATH=$(go env GOPATH)/bin:$PATH

generate:
	head -c 32 /dev/urandom | shasum --a 256

