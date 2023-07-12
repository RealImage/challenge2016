package datasouce

import (
	"sync"

	"chng2016/pkg/models"
)

// Cache ...
type Cache interface {
	SetCache(key string, value *models.Distributor)
	GetCache(key string) *models.Distributor
}

// CacheClient ...
type CacheClient struct {
	db map[string]*models.Distributor
	sync.RWMutex
}

// NewCacheClient ...
func NewCacheClient() *CacheClient {
	return &CacheClient{db: make(map[string]*models.Distributor)}
}
