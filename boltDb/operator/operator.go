package operator

import (

	"bytes"
	"encoding/gob"
	"execlt1/boltDb/gobCode"
	"execlt1/boltDb/student"
	"fmt"
	"go.etcd.io/bbolt"
	"path"
	"time"
)

type BoltDB struct {
	Prefix string
	dbName string
	db     *bbolt.DB
}

var (
	ErrNotExistsKey    = fmt.Errorf("key is not exists")
	ErrNotExistsBucket = fmt.Errorf("bucket is not exists")
	StudentHeight      = []byte("student_height")
)

func New(Prefix, dbName string) (*BoltDB, error) {
	db, err := bbolt.Open(dbName, 0600, &bbolt.Options{
		Timeout: time.Second * 1,
	})
	if err != nil {
		return nil, err
	}
	c := &BoltDB{
		Prefix: Prefix,
		dbName: dbName,
		db:     db,
	}
	gob.Register(student.Student{})
	return c, nil
}

func initBucket(db *bbolt.DB) error {
	if err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(StudentHeight)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *BoltDB) Sync() error {
	return c.db.Sync()
}

func (c *BoltDB) Close() error {
	if err := c.Sync(); err != nil {
		return err
	}
	if err := c.db.Close(); err != nil {
		return err
	}
	return nil
}

func (c *BoltDB) IsExistsKey(key string, bucketName string) bool {
	if err := c.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrNotExistsBucket
		}

		prefix := path.Join(c.Prefix, key)
		data := bucket.Get([]byte(prefix))
		if data == nil {
			return ErrNotExistsKey
		}
		return nil
	}); err != nil {
		return false
	}
	return true
}

func (c *BoltDB) Get(key string, into interface{}, bucketName string) (interface{}, error) {
	var data []byte
	if err := c.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrNotExistsBucket
		}
		prefix := path.Join(c.Prefix, key)
		data = bucket.Get([]byte(prefix))
		if data == nil {
			return ErrNotExistsKey
		}
		return nil
	}); err != nil {
		return nil, err
	}
	err := gobCode.Decode(data, &into)
	if err != nil {
		return nil, err
	}
	return into, nil
}

type DeletionFunc func(obj interface{}) bool

func (c *BoltDB) BatchDeletion(key string, into interface{}, fn DeletionFunc, bucketName string) error {
	if err := c.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrNotExistsBucket
		}
		prefix := []byte(path.Join(c.Prefix, key))

		cur := bucket.Cursor()
		for k, v := cur.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cur.Next() {
			if err := gobCode.Decode(v, &into); err == nil && fn(into) {
				_ = bucket.Delete(k)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

type ConditionFunc func(obj interface{}) bool

func (c *BoltDB) ConditionalList(key string, into interface{}, fn ConditionFunc, bucketName string) ([]interface{}, error) {
	list := make([]interface{}, 0)
	if err := c.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrNotExistsBucket
		}

		prefix := []byte(path.Join(c.Prefix, key))
		cur := bucket.Cursor()

		for k, v := cur.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = cur.Next() {
			err := gobCode.Decode(v, &into)
			if err != nil {
				return err
			}
			if fn(into) {
				list = append(list, into)
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return list, nil
}

func (c *BoltDB) Put(key string, obj interface{}, bucketName string) error {
	data, err := gobCode.Encode(obj)
	if err != nil {
		return err
	}
	if err := c.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		prefix := []byte(path.Join(c.Prefix, key))
		return bucket.Put(prefix, data)
	}); err != nil {
		return err
	}
	return nil
}

func (c *BoltDB) List(into interface{}, bucketName string) (map[string]interface{}, error) {
	dataMap := make(map[string]interface{})
	if err := c.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrNotExistsBucket
		}
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			gobCode.Decode(v, &into)
			dataMap[string(k)] = into
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return dataMap, nil
}

func (c *BoltDB) DeleteBucket(bucketName string) error {
	if err := c.db.Update(func(tx *bbolt.Tx) error {
		err := tx.DeleteBucket([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *BoltDB) Delete(key string, bucketName string) error {
	if err := c.db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return ErrNotExistsBucket
		}
		prefix := []byte(path.Join(c.Prefix, key))
		return bucket.Delete(prefix)
	}); err != nil {
		return err
	}
	return nil
}
