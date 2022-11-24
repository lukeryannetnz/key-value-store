package store

type LogItem struct {
	key    string
	value  string
	delete bool
}

type KeyValueStore struct {
	items        map[string]string
	pendingItems []LogItem
}

func (k *KeyValueStore) Set(key string, value string) {
	log := LogItem{key, value, false}

	k.pendingItems = append(k.pendingItems, log)
}

func (k *KeyValueStore) Get(key string) string {
	return k.items[key]
}

func (k *KeyValueStore) Delete(key string) {
	log := LogItem{key, "", true}

	k.pendingItems = append(k.pendingItems, log)
}

func (k *KeyValueStore) Begin() {
	k.items = make(map[string]string)
}

func (k *KeyValueStore) Commit() {
	for _, item := range k.pendingItems {
		if item.delete {
			delete(k.items, item.key)
		} else {
			k.items[item.key] = item.value
		}
	}
}

func (k *KeyValueStore) Rollback() {
	k.pendingItems = make([]LogItem, 0)
}
