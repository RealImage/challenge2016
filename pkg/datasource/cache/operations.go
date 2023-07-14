package datasouce

import (
	"strings"

	"chng2016/pkg/models"
)

// SetCache ...
func (c *CacheClient) SetCache(key string, value *models.Distributor) {
	c.Lock()
	defer c.Unlock()
	c.db[strings.ToLower(key)] = value
}

// GetCache ...
func (c *CacheClient) GetCache(key string) *models.Distributor {
	c.RLock()
	defer c.RUnlock()
	if val, ok := c.db[strings.ToLower(key)]; ok {
		return val
	}

	return nil
}
