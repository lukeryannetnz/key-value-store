package store

type KeyValueStore struct {
}

func (k *KeyValueStore) Set(key string, value string) {

}

func (k *KeyValueStore) Get(key string) string {
	return "test"
}

func (k *KeyValueStore) Delete(key string) {

}

func (k *KeyValueStore) Begin() {

}

func (k *KeyValueStore) Commit() {

}

func (k *KeyValueStore) Rollback() {

}
