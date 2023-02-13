package bitcoin

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	cl, err := Connect("locahost:8333", "user", "password")
	fmt.Println(err)
	fmt.Println(cl)
}
