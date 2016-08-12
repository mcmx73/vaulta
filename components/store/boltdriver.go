package store

import (
	"reflect"
	"github.com/boltdb/bolt"
	"github.com/DeepForestTeam/mobiussign/components/config"
	"github.com/DeepForestTeam/mobiussign/components/log"
	"github.com/mcmx73/vaulta/components/store/codecs/json"
	"github.com/mcmx73/vaulta/components/store/codecs/bson"
	"sync"
	"bytes"
	"encoding/binary"
)

type EncodeDecoder interface {
	Encode(v interface{}) ([]byte, error)
	Decode(b []byte, v interface{}) error
}

type BoltDriver struct {
	db        *bolt.DB
	connected bool
	Codec     EncodeDecoder
	mux       map[string]*sync.Mutex
}

type BoltIndex struct {
	ID  uint64
	Key string
}

const (
	indexPostfix = "_index"
)

type StorageObject struct {
	MainKey string `json:"primary_key"`
	Id      int64  `json:"id"`
	Data    interface{}
}

func (this *BoltDriver)Connect() (err error) {
	db_name, err := config.GetString("BOLT_DB")
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Debug("BoltDriver/Warper: Open BoltDB Storage, db file path:", db_name)
	this.db, err = bolt.Open(db_name, 0600, nil)
	if err != nil {
		log.Error("Can notopen BoltDB file:", err)
	}
	this.mux = make(map[string]*sync.Mutex)
	this.setDefaultCodec()
	return
}
func (this *BoltDriver)Close() {
	this.lockAll()
	defer this.unlockAll()
	this.db.Close()
}

func (this *BoltDriver)Set(bucket_name, key string, data interface{}) (id int64, err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	ref := reflect.ValueOf(data)
	if !ref.IsValid() || ref.Kind() != reflect.Ptr {
		return 0, ErrPtrNeeded
	}
	//Check is key exist:
	is_key_exist, err := this.IsKeyExist(bucket_name, key)
	if err != nil {
		log.Error("Can not check is key exist:", err)
		return
	}
	if is_key_exist {
		err = this.db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
			if err != nil {
				return err
			}
			store_object := StorageObject{}
			val := bucket.Get([]byte(key))
			err = this.Codec.Decode(val, &store_object)
			if err != nil {
				log.Error("Can not load store object:", err)
				return err
			}
			store_object.Data = data
			value, err := this.Codec.Encode(store_object)
			if err != nil {
				log.Error(err)
				return err
			}
			err = bucket.Put([]byte(key), value)
			if err != nil {
				return err
			}
			id = int64(store_object.Id)
			return nil
		})
		return
	} else {
		err = this.db.Update(func(tx *bolt.Tx) error {
			bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name))
			if err != nil {
				return err
			}
			store_object := StorageObject{}
			store_object.MainKey = key
			store_object.Data = data
			num, err := bucket.NextSequence()
			if err != nil {
				log.Error("Can not get next ID Value")
				return ErrInvalidIndex
			}
			store_object.Id = int64(num)
			value, err := this.Codec.Encode(store_object)
			if err != nil {
				log.Error(err)
				return err
			}
			err = bucket.Put([]byte(key), value)
			if err != nil {
				log.Error(err)
				return err
			}
			id = int64(num)
			//Create Index
			index_bucket, err := tx.CreateBucketIfNotExists([]byte(bucket_name + indexPostfix))
			if err != nil {
				return err
			}
			id := uintToBytes(num)
			err = index_bucket.Put(id, []byte(key))
			if err != nil {
				return err
			}
			return nil
		})
		return
	}

}

func (this *BoltDriver)Get(bucket_name, key string, data interface{}) (err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	ref := reflect.ValueOf(data)
	if !ref.IsValid() || ref.Kind() != reflect.Ptr {
		return ErrPtrNeeded
	}
	err = this.db.View(func(tx *bolt.Tx) error {
		bucket, err := this.getBucket(tx, bucket_name)
		if err != nil {
			return err
		}
		val := bucket.Get([]byte(key))
		if len(val) == 0 {
			return ErrKeyNotFound
		}
		store_object := StorageObject{}
		store_object.Data = data
		err = this.Codec.Decode(val, &store_object)
		if err == nil {
			data = store_object.Data
		}
		if err != nil {
			log.Error(err)
		}
		return err
	})
	return
}

func (this *BoltDriver)Count(bucket_name string) (count int64, err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	err = this.db.View(func(tx *bolt.Tx) error {
		bucket, err := this.getBucket(tx, bucket_name)
		if err != nil {
			return err
		}
		stats := bucket.Stats()
		count = int64(stats.KeyN)
		return err
	})
	return
}

func (this *BoltDriver)Last(bucket_name string, data interface{}) (key string, err error) {
	this.lockBucket(bucket_name)
	defer this.unlockBucket(bucket_name)
	ref := reflect.ValueOf(data)
	if !ref.IsValid() || ref.Kind() != reflect.Ptr {
		return key, ErrPtrNeeded
	}
	err = this.db.View(func(tx *bolt.Tx) error {
		index_bucket, err := this.getBucket(tx, bucket_name + indexPostfix)
		if err != nil {
			return err
		}
		_, index_val := index_bucket.Cursor().Last()
		if index_val == nil {
			return ErrNotIndexed
		}
		key = string(index_val)
		if err != nil {
			log.Error(err)
			return err
		}
		bucket, err := this.getBucket(tx, bucket_name)
		val := bucket.Get([]byte(key))
		if len(val) == 0 {
			return ErrKeyNotFound
		}
		store_object := StorageObject{}
		store_object.Data = data
		err = this.Codec.Decode(val, &store_object)
		if err == nil {
			data = store_object.Data
		}
		if err != nil {
			log.Error(err)
		}
		return err
	})

	return
}

func (this *BoltDriver)IsKeyExist(bucket_name string, key string) (exist bool, err error) {
	err = this.db.View(func(tx *bolt.Tx) error {
		bucket, err := this.getBucket(tx, bucket_name)
		if err != nil {
			return err
		}
		val := bucket.Get([]byte(key))
		if len(val) == 0 {
			return ErrKeyNotFound
		}
		return err
	})
	if err == nil {
		exist = true
	}
	if err == ErrKeyNotFound || err == ErrSectionNotFound {
		err = nil
	}
	return
}

func (this *BoltDriver)getBucket(tx *bolt.Tx, bucket_name string) (bucket *bolt.Bucket, err error) {
	bucket = tx.Bucket([]byte(bucket_name))
	if bucket == nil {
		return bucket, ErrSectionNotFound
	}
	return
}

func (this *BoltDriver)setCodec(name string) {
	switch name {
	case "json":
		this.Codec = json.Codec
	case "bson":
		this.Codec = bson.Codec
	}
}
func (this *BoltDriver)setDefaultCodec() {
	this.setCodec("json")
}
func (this *BoltDriver)lockBucket(bucket_name string) {
	if mux, ok := this.mux[bucket_name]; ok {
		mux.Lock()
	} else {
		mux = new(sync.Mutex)
		mux.Lock()
		this.mux[bucket_name] = mux
	}
}
func (this *BoltDriver)unlockBucket(bucket_name string) {
	if mux, ok := this.mux[bucket_name]; ok {
		mux.Unlock()
	}
}
func (this *BoltDriver)lockAll() {
	if len(this.mux) != 0 {
		for index, _ := range this.mux {
			this.mux[index].Lock()
		}
	}
}
func (this *BoltDriver)unlockAll() {
	if len(this.mux) != 0 {
		for index, _ := range this.mux {
			this.mux[index].Unlock()
		}
	}
}

func init() {

}
func uintToBytes(num uint64) []byte {
	buf := new(bytes.Buffer)
	_ = binary.Write(buf, binary.BigEndian, num)
	return buf.Bytes()
}


