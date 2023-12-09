package pkg

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type DBops interface {
	Commit(ctx context.Context, path string) error
	Rollback(ctx context.Context) error
	Begin(ctx context.Context) error
	Get(path string) (any, error)
	Set(key string, value any)
}
type recordData struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
type Database struct {
	baseFolder  string
	transaction map[string]interface{}
	inTx        bool
	txmutex     sync.Mutex
}

func NewDatabase(ctx context.Context, foldername string) (*Database, error) {
	if err := os.MkdirAll(foldername, 0755); err != nil {
		return nil, fmt.Errorf("unable to create database directory: %v", err)
	}
	return &Database{
		baseFolder: foldername,
	}, nil
}

func (db *Database) Commit(ctx context.Context, path string) error {
	db.txmutex.Lock()
	defer db.txmutex.Unlock()

	if !db.inTx {
		return errors.New("no active transaction")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	parts := strings.Split(path, "/")
	// we already checked in Parsekey
	keyID, _ := strconv.Atoi(parts[len(parts)-1])
	hashValue := keyID / 10000
	pathTableFolder := db.baseFolder + "/" + parts[0]

	if _, err := os.Stat(pathTableFolder); os.IsNotExist(err) {
		err := os.MkdirAll(pathTableFolder, 0755)
		if err != nil {
			return fmt.Errorf("Failed to create table folder directory: %v", err)
		}
	}
	file, err := os.OpenFile(pathTableFolder+"/"+strconv.Itoa(hashValue), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to Create file or write file %v", err)
	}
	defer file.Close()
	for key, value := range db.transaction {
		record := recordData{Key: key, Value: value}
		data, err := json.Marshal(record)
		if err != nil {
			log.Printf("Failed to Unmarshal key and value")
		}
		if _, err := file.Write(append(data, '\n')); err != nil {
			return err
		}
	}

	db.inTx = false
	// Clear the transaction map
	db.transaction = nil

	return nil
}
func (db *Database) Rollback(ctx context.Context) error {
	db.txmutex.Lock()
	defer db.txmutex.Unlock()
	if !db.inTx {
		return errors.New("no active transaction")
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	db.inTx = false
	db.transaction = nil // Discard the transaction map

	return nil
}

func (db *Database) Begin(ctx context.Context) error {
	db.txmutex.Lock()
	defer db.txmutex.Unlock()
	if db.inTx {
		return fmt.Errorf("Transaction already in progress.")
	}
	db.transaction = make(map[string]interface{})
	db.inTx = true
	return nil
}
func (db *Database) Get(path string) (any, error) {
	parts := strings.Split(path, "/")
	// we already checked in Parsekey
	keyID, _ := strconv.Atoi(parts[len(parts)-1])
	hashValue := keyID / 10000
	file, err := os.Open(db.baseFolder + "/" + parts[0] + "/" + strconv.Itoa(hashValue))
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("table folder %s does not exist", db.baseFolder+"/"+parts[0]+"/"+strconv.Itoa(hashValue)+"idx")
		}
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var record recordData
		line := scanner.Bytes()
		if err := json.Unmarshal(line, &record); err != nil {
			log.Printf("Failed to Unmarshal key and value")
		}
		if record.Key == parts[1] {
			return record.Value, nil // Key found, return the value.
		}
	}

	return nil, fmt.Errorf("Not Found")
}

func (db *Database) Set(path string, value any) {
	db.txmutex.Lock()
	defer db.txmutex.Unlock()
	parts := strings.Split(path, "/")
	if db.inTx {
		db.transaction[parts[len(parts)-1]] = value
	}
}
