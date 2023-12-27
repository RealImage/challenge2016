package store

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/challenge2016/model"
)

// LocalStorage is a simple file-based storage
type LocalStorage struct {
	file *os.File
	mu   sync.RWMutex
	Data map[string]model.Permissions
}

// NewLocalStorage creates a new LocalStorage instance
func NewLocalStorage(file *os.File) *LocalStorage {
	return &LocalStorage{
		file: file,
		Data: make(map[string]model.Permissions),
	}
}

// SaveData saves data to the local storage file
func (ls *LocalStorage) SaveData() error {
	ls.mu.RLock()
	defer ls.mu.RUnlock()

	dataJSON, err := json.Marshal(ls.Data)
	if err != nil {
		return err
	}
	err = ls.file.Truncate(0)
	_, err = ls.file.Seek(0, 0)
	_, err = ls.file.WriteString(string(dataJSON))
	return err
}

// LoadData loads data from the local storage file
func (ls *LocalStorage) LoadData() error {
	ls.mu.Lock()
	defer ls.mu.Unlock()

	fileContent, err := io.ReadAll(ls.file)
	if err != nil {
		return err
	}

	_ = json.Unmarshal(fileContent, &ls.Data)

	return nil
}

// GetPermissions retrieves permissions for a distributor
func (ls *LocalStorage) GetPermissions(distributorKey string) (model.Permissions, bool) {
	ls.mu.RLock()
	defer ls.mu.RUnlock()
	permissions, ok := ls.Data[distributorKey]
	return permissions, ok
}
