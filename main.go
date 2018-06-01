/*
nist-password-validator will detect passwords that do not meet the following requirements:

1. Minimum of 8 characters
1. Maximum of 64 characters
1. Only contain ASCII characters
1. Not be a common password based on a supplied common password file

The program will take a list of newline-delimited passwords from STDIN.
It will check each of these passwords against the above criteria and output any passwords that fail
to meet the criteria along with the failure reason.

A filename which contains the list of newline-delimited common passwords should be supplied
as the first parameter of the program.

	# Run with a single password
	echo "MyPassword" | nist-password-validator myCommonPasswordList.txt

	# Run with a file of passwords
	cat "myPasswordFile.txt" | nist-password-validator myCommonPasswordList.txt

Based on https://gist.github.com/buckhx/45a54e75f7a9484ca9d69f699b929eca
*/
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/caitlin615/nist-password-validator/password"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Common password file must be supplied as the first argument")
		os.Exit(1)
	}

	filename := os.Args[1]
	commonPasswordFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	validator := password.NewValidator(true, 8, 64)
	err = validator.AddCommonPasswords(commonPasswordFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Read the passwords from stdin
	stdin := bufio.NewScanner(os.Stdin)
	// TODO: What happens if nothing is provided to stdin?
	for stdin.Scan() {
		text := stdin.Text()
		if len(text) == 0 {
			continue
		}

		if err := validator.ValidatePassword(text); err != nil {
			if err == password.ErrInvalidCharacters {
				fmt.Printf("*** -> Error: %s\n", err)
			} else {
				fmt.Printf("%s -> Error: %s\n", text, err)
			}
		}
	}

	if err := stdin.Err(); err != nil {
		fmt.Printf("Error loading passwords from stdin: %s", err)
		os.Exit(1)
	}
}
