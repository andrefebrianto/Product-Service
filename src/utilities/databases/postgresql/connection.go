package postgresql

import (
	"errors"

	"github.com/go-pg/pg/v10"
)

type dbConnection struct {
	id string
	db *pg.DB
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
		opt, err := pg.ParseURL(uri)
		if err != nil {
			panic(err)
		}
		newConnection.db = pg.Connect(opt)

		connectionPools.connections = append(connectionPools.connections, newConnection)
	}
}

//GetConnection ...
func GetConnection(id string) (*pg.DB, error) {
	for _, connection := range connectionPools.connections {
		if connection.id == id {
			return connection.db, nil
		}
	}

	return nil, errors.New("Connection not found")
}
