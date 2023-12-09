package database

import (
	"RedisDurabilityDB/pkg"
	"context"
	"fmt"
	"time"
)

type ClientDatabase struct {
	db pkg.DBops
}

func NewClientDatabase(db pkg.DBops) *ClientDatabase {
	return &ClientDatabase{db: db}
}
func (c *ClientDatabase) Set(ctx context.Context, key string, value any, expiration time.Duration) error {

	// Begin a transaction
	err := c.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer c.db.Rollback(ctx) // This is safe to call even if the transaction is already committed.

	c.db.Set(key, value)

	return c.db.Commit(ctx, key)
}

func (c *ClientDatabase) Get(ctx context.Context, key string) (any, error) {
	value, err := c.db.Get(key)
	if err != nil {
		return nil, fmt.Errorf("no product found with id: %s", key)
	}
	return value, nil
}
