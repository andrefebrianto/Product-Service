package postgresql

import (
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type dbConnection struct {
	id string
	db interface{}
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
		newConnection.db = pg.Connect(&pg.Options{
			Addr: uri,
			User: "postgres",
		})

		connectionPools.connections = append(connectionPools.connections, newConnection)
	}
	fmt.Println(connectionPools)
}

//GetConnection ...
func GetConnection(id string) (interface{}, error) {
	fmt.Println(connectionPools)
	for _, connection := range connectionPools.connections {
		if connection.id == id {
			return connection.db, nil
		}
	}

	return nil, errors.New("Connection not found")
}
