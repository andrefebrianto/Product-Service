package postgresql

import (
	"context"
)

//Client ...
type Client struct {
	connection interface{}
}

//Insert ...
func Insert(context context.Context, query string, args ...interface{}) error {
	return nil
}

//Update ...
func Update(context context.Context, query string, args ...interface{}) error {
	return nil
}

//Upsert ...
func Upsert(context context.Context, query string, args ...interface{}) error {
	return nil
}

//Get ...
func Get(context context.Context, query string, args ...interface{}) ([]interface{}, error) {
	return nil, nil
}

//Delete ...
func Delete(context context.Context, query string, args ...interface{}) error {
	return nil
}
