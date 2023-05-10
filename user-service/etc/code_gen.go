package etc

import (
	"fmt"
	"io"
	"crypto/rand"
)

var (
	// table for code generator
	table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

func GenerateCode(max int) string {
	b := make([]byte, max)
	fmt.Println(max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		fmt.Println(err, "err ++++")
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}	

	return string(b)
}
