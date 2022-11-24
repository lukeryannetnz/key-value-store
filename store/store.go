package store

// an item in the transaction log
type logItem struct {
	key    string
	value  string
	delete bool
}

// the key value store, currently only supports a single transaction
type keyValueStore struct {
	items        map[string]string
	pendingItems []logItem
}

func NewKeyValueStore() keyValueStore {
	k := keyValueStore{}
	k.items = make(map[string]string)
	k.pendingItems = make([]logItem, 0)

	return k
}

func (k *keyValueStore) Set(key string, value string) {
	log := logItem{key, value, false}

	k.pendingItems = append(k.pendingItems, log)
}

func (k *keyValueStore) Get(key string) string {
	return k.items[key]
}

func (k *keyValueStore) Delete(key string) {
	log := logItem{key, "", true}

	k.pendingItems = append(k.pendingItems, log)
}

func (k *keyValueStore) Begin() {
}

func (k *keyValueStore) Commit() {
	for _, item := range k.pendingItems {
		if item.delete {
			delete(k.items, item.key)
		} else {
			k.items[item.key] = item.value
		}
	}

	k.pendingItems = make([]logItem, 0)
}

func (k *keyValueStore) Rollback() {
	k.pendingItems = make([]logItem, 0)
}
