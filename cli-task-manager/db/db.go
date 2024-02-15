package db

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type Task struct {
	Key   int
	Value string
}

var BoltDB *bolt.DB

const bucketName = "Tasks"

func InitDB(pathToFile string) {
	var err error
	BoltDB, err = bolt.Open("./tasks.db", 6060, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}

}

func getNextSequence(bucket *bolt.Bucket) (uint64, error) {
	id, err := bucket.NextSequence()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func WriteToDB(value []byte) error {

	err := BoltDB.Update(func(tx *bolt.Tx) error {

		b, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		id, err := getNextSequence(b)
		if err != nil {
			return err
		}
		return b.Put(itob(int(id)), value)
	})

	if err != nil {
		return err
	}
	return nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}

func RemoveFromDB(key int) error {
	err := BoltDB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		return b.Delete(itob(key))
	})
	if err != nil {
		return err
	}
	return nil
}

func ReadFromDB() ([]Task, error) {
	var tasks []Task
	err := BoltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("get bucket: FAILED")
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, Task{Key: btoi(k), Value: string(v)})
		}
		return nil

	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
