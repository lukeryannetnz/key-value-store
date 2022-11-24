package store

import "testing"

func TestSetHappyPath(t *testing.T) {
	store := KeyValueStore{}

	store.Begin()

	store.Set("foo", "bar")

	value := store.Get("foo")

	if value != "" {
		t.Fatal("shouldn't be able to get something which isn't committed")
	}

	store.Commit()

	value = store.Get("foo")

	if value != "bar" {
		t.Fatal("commit didn't stick")
	}
}
