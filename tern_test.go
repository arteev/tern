package tern_test

import (
	"testing"

	"github.com/arteev/tern"
)

func TestOp(t *testing.T) {
	res := tern.Op(false, "Yes", "No")
	if res != "No" {
		t.Errorf("Excepted No,but %v", res)
	}

	res = tern.Op(true, "Yes", "No")
	if res != "Yes" {
		t.Errorf("Excepted Yes,but %v", res)
	}
}
