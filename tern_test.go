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
	for n:=0; n< b.N; n ++ {
		tern.Op(true, "Yes", "No")
	}
}

func doTestdecode(t *testing.T,vals map[int] interface{}, args... interface{}) {
	for key,pair := range vals {
		res := tern.Decode(key,args[0],args[1],args[2:]...)
		if res!=pair {
			t.Errorf("Excepted %v, got %v",pair,res)
		}
	}
}

func TestDecode(t *testing.T) {
	doTestdecode(t,map[int] interface{} {
		1:"one",
		0:nil,
	},1,"one")
	doTestdecode(t,map[int] interface{} {
		1:nil,
		0:nil,
	},2,"two")
	doTestdecode(t,map[int] interface{} {
		1:"one",
		2:"two",
		0:nil,
		100:nil,
	},1,"one",2,"two")
}

func BenchmarkDecode(b *testing.B){
	for n := 0; n < b.N; n++ {
		tern.Decode("test","test2","value3","test3","value3","test","tryvalue","default")
	}
}