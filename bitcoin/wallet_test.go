package bitcoin

import (
	"fmt"
	"testing"
)

func TestNewAccount(t *testing.T) {
	btca := NewAccount()
	fmt.Println(btca)
}
