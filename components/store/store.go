package store

import (
	"github.com/DeepForestTeam/mobiussign/components/log"
)

type GlobalStore struct {
	driver StorageDriver
}

type StorageDriver interface {
	Connect() error
	Close()
	Set(section, key string, data interface{}) (int64, error)
	Get(section, key string, data interface{}) error
	Count(section string) (int64, error)
	Last(section string, data interface{}) (string, error)
	IsKeyExist(section string, key string) (bool, error)
	GetRandomBlock(section string, data interface{}) (string, error)
}

var storage GlobalStore

func init() {
	log.Info("* Init store")
	storage.driver = &BoltDriver{}
}

func ConnectDB() (err error) {
	err = storage.driver.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Debug("Connect Driver Storage")
	return
}

func Set(model_name, key string, object interface{}) (id int64, err error) {
	return storage.driver.Set(model_name, key, object)
}
func Get(model_name, key string, object interface{}) (err error) {
	return storage.driver.Get(model_name, key, object)
}
func Last(model_name string, object interface{}) (key string, err error) {
	return storage.driver.Last(model_name, object)
}

func IsKeyExist(model_name string, key string) (bool, error) {
	return storage.driver.IsKeyExist(model_name, key)
}

func Count(model_name string) (count int64, err error) {
	return storage.driver.Count(model_name)
}

func GetRandomBlock(model_name string, object interface{}) (key string, err error) {
	return storage.driver.GetRandomBlock(model_name, object)
}