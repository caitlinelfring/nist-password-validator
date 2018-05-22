// Based on https://gist.github.com/buckhx/45a54e75f7a9484ca9d69f699b929eca

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/caitlin615/nist-password-validator/password"
)

func main() {
	// Read the passwords from stdin
	stdin := bufio.NewScanner(os.Stdin)

	inputPasswords := []password.Password{}
	// TODO: What happens if nothing is provided to stdin?
	for stdin.Scan() {
		text := stdin.Text()
		if len(text) != 0 {
			inputPasswords = append(inputPasswords, password.Password(text))
		}
	}
	if err := stdin.Err(); err != nil {
		fmt.Printf("Error loading passwords from stdin: %s", err)
		os.Exit(1)
	}

	if len(os.Args) <= 1 {
		fmt.Println("Common password file must be supplied as the first argument")
		os.Exit(1)
	}

	filename := os.Args[1]
	commonPasswords, err := password.NewCommonList(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Check each password
	for _, pass := range inputPasswords {
		if err := pass.CheckValidity(commonPasswords); err != nil {
			if err == password.ErrInvalidCharacters {
				fmt.Printf("*** -> Error: %s\n", err)
			} else {
				fmt.Printf("%s -> Error: %s\n", pass, err)
			}
		}
	}
}
