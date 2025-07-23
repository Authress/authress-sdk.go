default: install

tidy:
	go mod tidy

install:
	go install .

test:
	go test -count=1 -parallel=4 ./...

integration:
	ENV=test go test -count=1 -parallel=4 -timeout 10m -v ./...

run-one-integration:
	ENV=test go test -count=1 -parallel=4 -timeout 10m -run "Test_authress_UsersApi" -v ./...