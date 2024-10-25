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
