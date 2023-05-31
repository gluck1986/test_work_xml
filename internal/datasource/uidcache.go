package datasource

// UIDCache implementation of IUIDCache
type UIDCache struct {
	storage map[int]struct{}
}

// NewUIDCache Constructor
func NewUIDCache() IUIDCache {
	return &UIDCache{storage: make(map[int]struct{})}
}

// Has in cache
func (t *UIDCache) Has(uid int) bool {
	_, ok := t.storage[uid]
	return ok
}

// Add to cache
func (t *UIDCache) Add(uid int) {
	t.storage[uid] = struct{}{}
}
