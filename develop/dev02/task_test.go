package main

import (
	"testing"
)

func TestFunction(t *testing.T) {

	expect := "aaaabccddddde"
	result := unpackString("a4bc2d5e")
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = "abcd"
	result = unpackString("abcd")
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = ""
	result = unpackString("45")
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = "qwe45"
	result = unpackString("qwe\\4\\5")
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = "qwe\\45"
	result = unpackString("qwe44444")
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}

	expect = "qwe\\\\5"
	result = unpackString("qwe\\\\\\\\\\")
	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
