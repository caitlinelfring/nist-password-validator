install:
	go build -o nist-password-validator

test: 10-million-password-list-top-1000000.txt
	go test -v -race ./...

test-coverage: 10-million-password-list-top-1000000.txt
	go test -v -covermode=count -coverprofile=coverage.out ./...

benchmark: 10-million-password-list-top-1000000.txt
	go test -v -bench . ./...

10-million-password-list-top-1000000.txt:
	curl -LO "https://github.com/danielmiessler/SecLists/raw/master/Passwords/Common-Credentials/$@"
