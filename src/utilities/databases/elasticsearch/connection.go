package elasticsearch

import (
	"errors"

	"github.com/elastic/go-elasticsearch/v7"
)

type dbConnection struct {
	id string
	db *elasticsearch.Client
}

//ConnectionPool ...
type ConnectionPool struct {
	connections []dbConnection
}

var (
	connectionPools ConnectionPool
)

//InitConnection ...
func InitConnection(configs []map[string]interface{}) {
	for _, config := range configs {
		newConnection := dbConnection{}
		id := config["id"].(string)
		newConnection.id = id
		uri := config["uri"].(string)
		addresses := []string{uri}
		cfg := elasticsearch.Config{Addresses: addresses}
		dbClient, err := elasticsearch.NewClient(cfg)
		if err != nil {
			panic(err)
		}

		newConnection.db = dbClient
		connectionPools.connections = append(connectionPools.connections, newConnection)
	}
}

//GetConnection ...
func GetConnection(id string) (*elasticsearch.Client, error) {
	for _, connection := range connectionPools.connections {
		if connection.id == id {
			return connection.db, nil
		}
	}

	return nil, errors.New("Connection not found")
}
