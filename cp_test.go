package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCp(t *testing.T) {
	f, _ := ioutil.TempFile(".", "")
	Copy(f.Name(), "xxx")
	_, err := os.Open("xxx")
	if err != nil {
		t.Error("Copy failed.")
	}
	os.Remove(f.Name())
	os.Remove("xxx")
}
