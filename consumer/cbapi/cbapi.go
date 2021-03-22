package cbapi

import (
	"errors"
	"time"

	"github.com/couchbase/gocb/v2"
)

const (
	KEY    = "COUNTER"
	SUBDOC = "counter"
)

var (
	coll  *gocb.Collection
	DB_IP string
)

type Doc struct {
	Counter int `json:"counter"`
}

func InitDB() error {
	opts := gocb.ClusterOptions{Username: "Administrator", Password: "password"}
	Cluster, err := gocb.Connect(DB_IP, opts)
	if err != nil {
		return err
	}
	bucket := Cluster.Bucket("ut")
	if err != nil {
		return err
	}
	coll = bucket.DefaultCollection()
	_, err = coll.Insert(KEY, &Doc{Counter: 0}, &gocb.InsertOptions{Timeout: time.Second * 5})
	if err != nil {
		if errors.Is(err, gocb.ErrDocumentExists) {
			return nil
		} else {
			return err
		}
	}
	return nil
}

func ClosrDBConnection() {
}

type DBConnection struct {
}

type doc struct {
	ID int `json:"counter"`
}

func (conn *DBConnection) FetchCounter() (int, error) {
	ops := make([]gocb.LookupInSpec, 1)
	ops[0] = gocb.GetSpec("counter", nil)

	result, err := coll.LookupIn(KEY, ops, &gocb.LookupInOptions{Timeout: 5 * time.Second})
	if err != nil {
		return 0, err
	}
	var counter int
	err = result.ContentAt(0, &counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (conn *DBConnection) UpdateCounter() error {
	ops := make([]gocb.MutateInSpec, 1)
	ops[0] = gocb.IncrementSpec("counter", 1, &gocb.CounterSpecOptions{CreatePath: true})

	_, err := coll.MutateIn(KEY, ops, &gocb.MutateInOptions{Timeout: 5 * time.Second})
	if err != nil {
		return err
	}

	return nil
}
