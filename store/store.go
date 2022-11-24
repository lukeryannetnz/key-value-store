package store

type LogItem struct {
	key    string
	value  string
	delete bool
}

type keyValueStore struct {
	items        map[string]string
	pendingItems []LogItem
}

func NewKeyValueStore() keyValueStore {
	k := keyValueStore{}
	k.items = make(map[string]string)
	k.pendingItems = make([]LogItem, 0)

	return k
}

func (k *keyValueStore) Set(key string, value string) {
	log := LogItem{key, value, false}

	k.pendingItems = append(k.pendingItems, log)
}

func (k *keyValueStore) Get(key string) string {
	return k.items[key]
}

func (k *keyValueStore) Delete(key string) {
	log := LogItem{key, "", true}

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

	k.pendingItems = make([]LogItem, 0)
}

func (k *keyValueStore) Rollback() {
	k.pendingItems = make([]LogItem, 0)
}
