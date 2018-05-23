package password

import (
	"math/rand"
	"os"
	"testing"
)

func BenchmarkCommonListMatches(b *testing.B) {
	var filename = "../10-million-password-list-top-1000000.txt"
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		b.Skipf("top 1million common passwords file (%s) doesn't exist. Try running `make 10-million-password-list-top-1000000.txt`.", filename)
	}
	commonList, err := NewCommonList(filename)
	if err != nil {
		b.Fatal(err)
	}

	// Typically a non-fixed seed should be used, such as time.Now().UnixNano().
	// Using a fixed seed will produce the same output on every run.
	r := rand.New(rand.NewSource(99))
	totalCommonPasswords := len(commonList.list)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		randCommon := commonList.list[r.Intn(totalCommonPasswords)]
		// These should all match since we're pulling it from the list of common passwords
		if !commonList.Matches(randCommon) {
			b.Fatalf("Common password match failed: %s", randCommon)
		}
	}
}
