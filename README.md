# NIST Password Validator
[![Go Report Card](https://goreportcard.com/badge/github.com/caitlin615/nist-password-validator)](https://goreportcard.com/report/github.com/caitlin615/nist-password-validator)
[![GoCover.io](https://gocover.io/_badge/github.com/caitlin615/nist-password-validator/password)](https://gocover.io/github.com/caitlin615/nist-password-validator/password)
[![GoDoc](https://godoc.org/github.com/caitlin615/nist-password-validator?status.svg)](https://godoc.org/github.com/caitlin615/nist-password-validator)

[NIST](https://www.nist.gov/) recently updates their [Digital Identity Guidelines](https://pages.nist.gov/800-63-3/) in June 2017.
The new guidelines specify general rules for handling the security of user supplied passwords.

This program will detect passwords that do not meet the following requirements:

1. Minimum of 8 characters
1. Maximum of 64 characters
1. Only contain ASCII characters
1. Not be a common password based on a supplied common password file

The program will take a list of newline-delimited passwords from STDIN.
It will check each of these passwords against the above criteria and output any passwords that fail
to meet the criteria along with the failure reason.

A filename which contains the list of common passwords should be supplied as the first parameter of the program.

# Installing and running

First, make sure you have the latest version of [Go](https://golang.org/doc/install) installed.

To install the binary:

```bash
go install github.com/caitlin615/nist-password-validator
```

To run it:
* `myCommonPasswordList.txt` is a newline-delimited file containing a list of common passwords.
You will need to supply this yourself, or [see here](#downloading-common-password-list)

```bash
# Run with a single password
echo "MyPassword" | nist-password-validator myCommonPasswordList.txt

# Run with a file of passwords
cat "myPasswordFile.txt" | nist-password-validator myCommonPasswordList.txt
```

# Building and running locally

```bash
# Download the repo
go get -d github.com/caitlin615/nist-password-validator
cd $GOPATH/src/nist-password-validator

# Download the common password list (see Makefile for more details)
make 10-million-password-list-top-1000000.txt

# Build the binary
go build -o nist-password-validator
```

Now you can run it!

```bash
# Run with a single password
echo "MyPassword" | ./nist-password-validator myCommonPasswordList.txt

# Run with a file of passwords
cat "myPasswordFile.txt" | ./nist-password-validator myCommonPasswordList.txt
```

# Expected output

The program will only output a list of passwords that fail the validation tests
and their reason for failure.

Here's an example:

```
abc123 -> Error: Too Short
password -> Error: Common Password
abcd -> Error: Too Short
asdlfdja;lsdijfa;sdfpoaisdufpoaisuf9oarpfauhrfiuehfpiudhfioufgoiudfhgoiudfgpdupodsifpuosiUFPAOSIDUFPAOSDIUFP -> Error: Too Long
*** -> Error: Invalid Characters
*** -> Error: Invalid Characters
*** -> Error: Invalid Characters
*** -> Error: Invalid Characters
```

## Downloading Common Password List

You can download a common password list using the following command:

```bash
curl -LO "https://github.com/danielmiessler/SecLists/raw/master/Passwords/Common-Credentials/10-million-password-list-top-1000000.txt"
# OR
make 10-million-password-list-top-1000000.txt
```


### Running Tests
```bash
make test
# OR
go test -v ./...
```
