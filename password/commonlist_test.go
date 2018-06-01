package password

import (
	"math/rand"
	"os"
	"testing"
)

func TestNewCommonPassword(t *testing.T) {
	tests := []struct {
		filename         string
		expectFileExists bool
	}{
		{
			filename:         "password_list_test.txt",
			expectFileExists: true,
		}, {
			filename:         "i_shouldn't_exists.txt",
			expectFileExists: false,
		},
	}

	for _, test := range tests {
		file, err := os.Open(test.filename)
		fileDoesntExists := err != nil && os.IsNotExist(err)
		if !fileDoesntExists != test.expectFileExists {
			t.Fatalf("%s expected: %t, got: %t", test.filename, test.expectFileExists, fileDoesntExists)
		}
		if !fileDoesntExists {
			_, err = NewCommonList(file)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

// TODO: Benchmark comparison of map vs list of strings
func BenchmarkCommonListMatches(b *testing.B) {
	var filename = "../10-million-password-list-top-1000000.txt"
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		b.Skipf("top 1million common passwords file (%s) doesn't exist. Try running `make 10-million-password-list-top-1000000.txt`.", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		b.Fatal(err)
	}
	commonList, err := NewCommonList(file)
	if err != nil {
		b.Fatal(err)
	}

	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	r := rand.New(rand.NewSource(99))
	totalCommonPasswords := len(commonList.mapping)

	// So the benchmark test can pick random items from the common password list,
	// we're putting them in slice of strings
	listing := make([]string, 0, totalCommonPasswords)
	for p := range commonList.mapping {
		listing = append(listing, p)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randCommon := listing[r.Intn(totalCommonPasswords)]
		// These should all match since we're pulling it from the list of common passwords
		if !commonList.Matches(randCommon) {
			b.Errorf("Common password match failed: %s", randCommon)
		}
	}
}
