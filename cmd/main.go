package main

import (
	"RedisDurabilityDB/internal/controller"
	"RedisDurabilityDB/internal/datasource/cache"
	"RedisDurabilityDB/internal/datasource/database"
	"RedisDurabilityDB/pkg"
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()

	databases, err := pkg.NewDatabase(ctx, "Persistent Database")
	if err != nil {
		log.Fatalf("Failed to connect Database %s", err)
	}
	// Initialize cache
	cacheInstance := pkg.NewCache()
	// datasource for cache
	dbClient := database.NewClientDatabase(databases)
	cacheClient := cache.NewClientCache(cacheInstance, dbClient)
	cacheController := controller.NewController(cacheClient)
	// Set data in cache

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the My database application.")
	fmt.Println("Operation Types: Insert and Get")
	fmt.Println("-> If you want to exit type => Exit")
	for {
		fmt.Println("-> Which operation?")
		fmt.Print("-> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("-> Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "" {
			fmt.Println("Please input operation type!")
			continue
		}

		args := strings.Fields(input)
		if len(args) > 1 {
			fmt.Println("-> Invalid Operation")
			continue
		}
		lowercaseWord := strings.ToLower(args[0])

		if lowercaseWord == "exit" {
			fmt.Println("-> Goodbye!")
			return
		}
		if lowercaseWord == "insert" || lowercaseWord == "get" {
			var key string
			for {
				fmt.Println("-> Key: table: [table_name], keyID: [key_id]")
				fmt.Print("-> ")
				key, err = reader.ReadString('\n')
				if err != nil {
					fmt.Println("-> Error reading input:", err)
					continue
				}
				key = strings.TrimSpace(key)
				if key == "" {
					fmt.Println("Please input key!")
					continue
				}
				break
			}
			if lowercaseWord == "insert" {
				var value string
				for {
					fmt.Println("-> Value:")
					fmt.Print("-> ")
					value, err = reader.ReadString('\n')
					if err != nil {
						fmt.Println("-> Error reading inout:", err)
						continue
					}
					value = strings.TrimSpace(value)
					if value == "" {
						fmt.Printf("Please input value for key %s: \n", key)
						continue
					}
					break
				}
				err = cacheController.Set(ctx, key, value, 1*time.Second)
				if err != nil {
					fmt.Printf("Failed to update cache database: %s", err)
					continue
				}
				continue
			}
			got, err := cacheController.Get(ctx, key)
			if err != nil {
				fmt.Printf("Failed to get value from cache: %s", err)
				continue
			}
			fmt.Println(got)
		}
	}

}
