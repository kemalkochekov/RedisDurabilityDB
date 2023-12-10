package database

import (
	"context"
	"fmt"
	"time"

	"RedisDurabilityDB/pkg"
)

type ClientDatabase struct {
	db pkg.DBops
}

func NewClientDatabase(db pkg.DBops) *ClientDatabase {
	return &ClientDatabase{db: db}
}

func (c *ClientDatabase) Set(ctx context.Context, key string, value any, _ time.Duration) error {
	// Begin a transaction
	err := c.db.Begin(ctx)
	if err != nil {
		return err
	}

	c.db.Set(key, value)

	err = c.db.Commit(ctx, key)
	if err != nil {
		if txErr := c.db.Rollback(ctx); txErr != nil {
			return txErr
		}

		return err
	}

	return nil
}

func (c *ClientDatabase) Get(_ context.Context, key string) (any, error) {
	value, err := c.db.Get(key)
	if err != nil {
		return nil, fmt.Errorf("no product found with id: %s", key)
	}

	return value, nil
}
