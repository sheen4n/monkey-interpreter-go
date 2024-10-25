package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("hello1.HashKey() not equal to hello2.HashKey()")
	}

	if diff1.HashKey() == diff2.HashKey() {
		t.Errorf("diff1.HashKey() equal to diff2.HashKey()")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("hello1.HashKey() equal to diff1.HashKey()")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("true1.HashKey() not equal to true2.HashKey()")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("false1.HashKey() not equal to false2.HashKey()")
	}

	if true1.HashKey() == false1.HashKey() {
		t.Errorf("true1.HashKey() equal to false1.HashKey()")
	}
}
