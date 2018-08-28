package database

import (
	"github.com/gobuffalo/pop"
)

var dbConnection pop.Connection

func init() {
	dbConnection, err := pop.Connect("db")
	if err != nil {
		panic(err)
	}
	// hack to skip unused var error
	_ = dbConnection
}
