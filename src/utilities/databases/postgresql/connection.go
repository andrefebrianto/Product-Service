package postgresql

import (
	"errors"

	"github.com/go-pg/pg/v10"
)

type dbConnection struct {
	id string
	db interface{}
}

//ConnectionPool ...
type ConnectionPool struct {
	Connections []dbConnection
}

//InitConnection ...
func InitConnection(configs []map[string]interface{}) ConnectionPool {
	connectionPool := ConnectionPool{}

	for _, config := range configs {
		newConnection := dbConnection{}
		id := config["id"].(string)
		newConnection.id = id
		uri := config["uri"].(string)
		newConnection.db = pg.Connect(&pg.Options{
			Addr: uri,
			User: "postgres",
		})

		connectionPool.Connections = append(connectionPool.Connections, newConnection)
	}

	return connectionPool
}

//GetConnection ...
func (connectionPool *ConnectionPool) GetConnection(id string) (interface{}, error) {
	for _, connection := range connectionPool.Connections {
		if connection.id == id {
			// context := context.Background()
			// if err := connection.db.Ping(context); err != nil {
			// 	return nil, err
			// }
			return connection.db, nil
		}
	}

	return nil, errors.New("Connection not found")
}
