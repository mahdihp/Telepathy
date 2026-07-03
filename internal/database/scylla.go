package database

import (
	"fmt"

	gocql "github.com/apache/cassandra-gocql-driver/v2"
	"github.com/mahdihp/telepathy/configs"
)

func NewScylla(cfg configs.Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.ScyllaHost)

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	// ایجاد Keyspace در صورت عدم وجود
	err = session.Query(`
		CREATE KEYSPACE IF NOT EXISTS ` + cfg.KeySpace +
		` WITH replication = {
			'class': 'NetworkTopologyStrategy',
			'replication_factor': 1
		};
	`).Exec()
	if err != nil {
		session.Close()
		fmt.Println(err)
		return nil, err
	}
	session.Close()
	cluster.Keyspace = cfg.KeySpace

	session, err = cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}
