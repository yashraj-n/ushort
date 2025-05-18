package repository

import (
	"sync"
	"time"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Links struct {
	Redirect  string    `json:"redirect"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	IpAddr    string    `json:"ip_addr"`
}

func (l Links) String() string {
	return "Links{Redirect: " + l.Redirect + ", Hash: " + l.Hash + ", CreatedAt: " + l.CreatedAt.String() + ", IpAddr: " + l.IpAddr + "}"
}

var lruSingleton *lru.Cache[string, Links]
var lruSingletoneLock sync.Mutex

// var singletoneLock sync.Mutex
var MAX_LRU_LEN = 100

func GetLruCache() *lru.Cache[string, Links] {
	if lruSingleton != nil {
		return lruSingleton
	}

	lruSingletoneLock.Lock()
	defer lruSingletoneLock.Unlock()

	var err error
	lruSingleton, err = lru.New[string, Links](MAX_LRU_LEN)
	if err != nil {
		return nil
	}
	return lruSingleton
}
