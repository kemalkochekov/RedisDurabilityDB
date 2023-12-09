package cache

import (
	"RedisDurabilityDB/internal/datasource"
	"RedisDurabilityDB/pkg"
	"context"
	"errors"
	"fmt"
	"time"
)

type ClientCache struct {
	connCache *pkg.Cache
	source    datasource.Datasource
}

func NewClientCache(connCache *pkg.Cache, source datasource.Datasource) *ClientCache {
	return &ClientCache{connCache: connCache, source: source}
}

func (c *ClientCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {

	// Call Set on the underlying cache implementation
	err := c.connCache.Set(ctx, key, value, expiration)
	if err != nil {
		return fmt.Errorf("Unable to set in cache")
	}
	err = c.source.Set(ctx, key, value, expiration)
	if err != nil {
		return errors.New("unable to Add into database")
	}
	return nil
}

func (c *ClientCache) Get(ctx context.Context, key string) (any, error) {
	// Try to get the value from the cache
	val, err := c.connCache.Get(ctx, key)
	if err != nil {
		dataDb, err := c.source.Get(ctx, key)
		if err != nil {
			return nil, fmt.Errorf("key not found in database")
		}
		err = c.connCache.Set(ctx, key, dataDb, 100*time.Second)
		if err != nil {
			return nil, err
		}
		return dataDb, nil
	}
	return val, nil
}
