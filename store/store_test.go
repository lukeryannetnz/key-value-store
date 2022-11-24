package store

import "testing"

func TestSetGet(t *testing.T) {
	store := NewKeyValueStore()

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

func TestMultipleCommits(t *testing.T) {
	store := NewKeyValueStore()

	store.Begin()

	store.Set("foo", "bar")

	store.Commit()

	store.Begin()

	store.Set("baz", "bing")

	store.Commit()

	value := store.Get("foo")

	if value != "bar" {
		t.Fatal("commit 1 didn't stick")
	}

	value2 := store.Get("baz")

	if value2 != "bing" {
		t.Fatal("commit 2 didn't stick")
	}
}

func TestDelete(t *testing.T) {
	store := NewKeyValueStore()

	store.Begin()

	store.Set("foo", "bar")
	store.Set("baz", "bing")
	store.Delete("foo")

	store.Commit()

	value := store.Get("foo")

	if value != "" {
		t.Fatal("delete didn't stick")
	}
}

func TestRollback(t *testing.T) {
	store := NewKeyValueStore()

	store.Begin()

	store.Set("foo", "bar")
	store.Set("baz", "bing")

	store.Rollback()

	value := store.Get("foo")

	if value != "" {
		t.Fatal("rollback didn't work")
	}
}
