package common

import (
	"fmt"
	"testing"
)

func TestValidAddress(t *testing.T) {
	var addr Address
	if addr != InvalidAddr {
		t.Fatal("invalid address")
	}
	fmt.Println(addr[:])
	fmt.Println(InvalidAddr[:])
}
