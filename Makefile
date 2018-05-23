install:
	go build -o nist-password-validator

test: 10-million-password-list-top-1000000.txt
	go test -v ./...

10-million-password-list-top-1000000.txt:
	curl -LO "https://github.com/danielmiessler/SecLists/raw/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt"
