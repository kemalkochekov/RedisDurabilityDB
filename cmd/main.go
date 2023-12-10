package main

import (
	"RedisDurabilityDB/internal/controller"
	"RedisDurabilityDB/internal/core"
	"RedisDurabilityDB/internal/datasource/cache"
	"RedisDurabilityDB/internal/datasource/database"
	"RedisDurabilityDB/pkg"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)

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
	cli := core.NewCLI(cacheController)
	cli.Run()
	fmt.Println("Welcome to the My database application.")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	go func() {
		<-quit
		// Perform graceful shut down
		stop()

		os.Exit(0)
	}()
}
