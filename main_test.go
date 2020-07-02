package main

import (
	"reflect"
	"testing"
)

func TestMain0(t *testing.T) {
	read0("testdata/main/0.txt")
	read0("testdata/main/1.txt")
	read0("testdata/main/2.txt")
	read0("testdata/main/3.txt")
}

func TestMain1(t *testing.T) {

	b, err := readTrue("testdata/main/0.txt")
	if !(b == false && err != nil && reflect.TypeOf(err).String() == `*os.PathError`) {
		t.Fatal("case 0")
	}

	b, err = readTrue("testdata/main/1.txt")
	if !(b == true && err == nil) {
		t.Fatal("case 1")
	}

	b, err = readTrue("testdata/main/2.txt")
	if !(b == false && err == nil) {
		t.Fatal("case 2")
	}

	b, err = readTrue("testdata/main/3.txt")
	if !(b == false && err != nil && reflect.TypeOf(err).String() == `*errors.errorString`) {
		t.Fatal("case 3")
	}
}
