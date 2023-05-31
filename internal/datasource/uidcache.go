package datasource

type UidCache struct {
	storage map[int]struct{}
}

func NewUidCache() IUidCache {
	return &UidCache{storage: make(map[int]struct{})}
}

func (t *UidCache) Has(uid int) bool {
	_, ok := t.storage[uid]
	return ok
}

func (t *UidCache) Add(uid int) {
	t.storage[uid] = struct{}{}
}
