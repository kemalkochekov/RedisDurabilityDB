package controller

import (
	"RedisDurabilityDB/internal/datasource"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const keyValue = 2

type Controller struct {
	source datasource.Datasource
}

func NewController(source datasource.Datasource) *Controller {
	return &Controller{source: source}
}

func (c *Controller) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	key, err := parsedKey(key)
	if err != nil {
		return fmt.Errorf("Failed to parse key %w", err)
	}

	return c.source.Set(ctx, key, value, expiration)
}

func (c *Controller) Get(ctx context.Context, key string) (any, error) {
	parsedPathID, err := parsedKey(key)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse key %w", err)
	}

	return c.source.Get(ctx, parsedPathID)
}

func parsedKey(key string) (string, error) {
	pairs := strings.Split(key, ",")
	parsedData := make(map[string]string)

	for _, pair := range pairs {
		pair := strings.Trim(pair, " ")
		if pair == "" {
			continue
		}

		parts := strings.Split(pair, ":")
		if len(parts) == keyValue {
			key := strings.Trim(parts[0], " ")
			value := strings.Trim(parts[1], " ")

			if value == "" || key == "" {
				return "", fmt.Errorf("Value or Key cannot be empty")
			}
			parsedData[key] = value

			continue
		}

		return "", fmt.Errorf("invalid input value, must be \"key: value\"")
	}

	keyID, found := parsedData["keyID"]
	if !found {
		return "", fmt.Errorf("keyID not found")
	}

	_, err := strconv.Atoi(keyID)
	if err != nil {
		return "", fmt.Errorf("keyID %s is not an integer", keyID)
	}

	tableName, found := parsedData["table"]
	if !found {
		return "", fmt.Errorf("Table not found")
	}

	return tableName + "/" + keyID, nil
}
