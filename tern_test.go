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

func BenchmarkOp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tern.Op(true, "Yes", "No")
	}
}

func doTestdecode(t *testing.T, vals map[int]interface{}, args ...interface{}) {
	for key, pair := range vals {
		res := tern.Decode(key, args[0], args[1], args[2:]...)
		if res != pair {
			t.Errorf("Excepted %v, got %v", pair, res)
		}
	}
}

func TestDecode(t *testing.T) {
	doTestdecode(t, map[int]interface{}{
		1: "one",
		0: nil,
	}, 1, "one")
	doTestdecode(t, map[int]interface{}{
		1: nil,
		0: nil,
	}, 2, "two")
	doTestdecode(t, map[int]interface{}{
		1:   "one",
		2:   "two",
		0:   nil,
		100: nil,
	}, 1, "one", 2, "two")

	doTestdecode(t, map[int]interface{}{
		1: nil,
		0: nil,
	}, 2, "two")

	doTestdecode(t, map[int]interface{}{
		1: nil,
	}, 2, "two", nil)
}

func BenchmarkDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tern.Decode("test", "test2", "value3", "test3", "value3", "test", "tryvalue", "default")
	}
}

func TestNvl(t *testing.T) {
	var val interface{}

	val = tern.Nvl("test", "def")
	if val != "test" {
		t.Errorf("Excepted %q,got %v", "test", val)
	}

	var nilval interface{}
	val = tern.Nvl(nilval, "default value")
	if val != "default value" {
		t.Errorf("Excepted %q,got %v", "default value", val)
	}

}

func TestCoalesce(t *testing.T) {
	var val interface{}
	var expr1 = "test"
	var expr2 interface{}

	val = tern.Coalesce(nil, nil)
	if val != nil {
		t.Errorf("Excepted nil,got %v", val)
	}
	val = tern.Coalesce(nil, nil, nil)
	if val != nil {
		t.Errorf("Excepted nil,got %v", val)
	}

	val = tern.Coalesce(expr1, nil)
	if val != expr1 {
		t.Errorf("Excepted %v,got %v", expr1, val)
	}

	val = tern.Coalesce(nil, expr1)
	if val != expr1 {
		t.Errorf("Excepted %v,got %v", expr1, val)
	}

	val = tern.Coalesce(nil, expr2, expr1)
	if val != expr1 {
		t.Errorf("Excepted %v,got %v", expr1, val)
	}

	val = tern.Coalesce(nil, nil, nil, nil, expr1, nil, nil)
	if val != expr1 {
		t.Errorf("Excepted %v,got %v", expr1, val)
	}

}
