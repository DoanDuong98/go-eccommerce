APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/
test:
	go test -v