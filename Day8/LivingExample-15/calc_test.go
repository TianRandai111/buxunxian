package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := add(2, 5)
	if r != 6 {
		t.Fatalf("add 2+6 != 6 err ,expect:%d,actual:%d", 6, r)
	}
	t.Logf("test add succ")
}

func TestSub(t *testing.T) {
	r := sub(5, 2)
	if r != 3 {
		t.Fatalf("add 2+6 != 6 err ,expect:%d,actual:%d", 3, r)
	}
	t.Logf("test sub succ")
}
