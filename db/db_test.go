package db

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	var testdb DB
	testdb, err := Open("test.db")
	defer testdb.Close()
	_ = testdb
	if err != nil {
		t.Errorf("something wrong happened: %s", err)
	}
}

func TestGet(t *testing.T) {
	testdb, _ := Open("test.db")
	// defers are last in first out
	defer os.Remove("test.db")
	defer testdb.Close()
	testdb.Put("mykey", "myvalue")
	v := testdb.Get("mykey")
	if v != "myvalue" {
		t.Errorf("expected: myvalue, actual: %v", v)
	}

	v = testdb.Get("doesntexist")
	if v != "" {
		t.Errorf("expected: <empty string>, actual: %v", v)
	}
}

func TestPut(t *testing.T) {
	testdb, _ := Open("test.db")
	// defers are last in first out
	defer os.Remove("test.db")
	defer testdb.Close()
	err := testdb.Put("mykey", "myvalue")
	if err != nil {
		t.Errorf("something wrong happened: %s", err)
	}
}
