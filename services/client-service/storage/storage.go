// Package that encapsulate storage for gin.Context connections.
package storage

import (
	"go-microservices/services/client-service/constants"
	"sync"
	"time"
)

var sg *Storage = nil

// Type that encapsulate logic for CLIENT connections storage.
type Storage struct {
	sync.Mutex
	Storage []*StorageItem
}

// Type that encapsulate logic for CLIENT.
type StorageItem struct {
	sync.Mutex
	UUID string
	Chan chan constants.InternalMessage
	Ago  int64
}

// Add CLIENT to STORAGE.
func (self *Storage) AddToStorage(UUID string, ch chan constants.InternalMessage) *StorageItem {
	self.Lock()
	defer self.Unlock()
	return self.addToStorage(UUID, ch)
}

// Add CLIENT to STORAGE.
func (self *Storage) addToStorage(UUID string, ch chan constants.InternalMessage) *StorageItem {
	removes := []string{}
	for _, item := range self.Storage {
		if time.Now().Unix()-item.Ago > constants.TTL {
			removes = append(removes, item.UUID)
		}
	}

	for _, uuid := range removes {
		self.removeFromStorage(uuid)
	}

	item := new(StorageItem)
	item.UUID = UUID
	item.Chan = ch
	item.Ago = time.Now().Unix()
	self.Storage = append(self.Storage, item)
	return item
}

// Remove CLIENT to STORAGE.
func (self *Storage) removeFromStorage(UUID string) {
	if len(self.Storage) > 0 {
		findIndex := -1
		for i, v := range self.Storage {
			if v.UUID == UUID {
				findIndex = i
				break
			}
		}
		self.Storage = append(self.Storage[:findIndex], self.Storage[findIndex+1:]...)
	}
}

// Remove CLIENT to STORAGE.
func (self *Storage) RemoveFromStorage(UUID string) {
	self.Lock()
	defer self.Unlock()
	self.removeFromStorage(UUID)
}

// Get conn by UUID.
func (self *Storage) GetChan(uuid string) chan constants.InternalMessage {
	self.Lock()
	defer self.Unlock()
	return self.getChan(uuid)
}

// Get conn by UUID.
func (self *Storage) getChan(uuid string) chan constants.InternalMessage {
	for _, item := range self.Storage {
		if (uuid == item.UUID && item.Chan != nil && time.Now().Unix()-item.Ago <= constants.TTL) {
			return item.Chan
		}
	}

	return nil
}

// Get STORAGE.
func GetStorage() *Storage {
	if sg == nil {
		sg = new(Storage)
		sg.Storage = []*StorageItem{}
	}
	return sg
}
