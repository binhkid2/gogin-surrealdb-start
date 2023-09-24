package database

import (
	util "github.com/binhkid2/gogin-surrealdb-start/Util"
	"github.com/surrealdb/surrealdb.go"
)

var db *surrealdb.DB

func init() {
	var err error
	db, err = surrealdb.New(util.GetConfig("database", "url"))
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": util.GetConfig("database", "user"),
		"pass": util.GetConfig("database", "pass"),
	}); err != nil {
		panic(err)
	}
}

func Connect() {
	var err error
	db, err = surrealdb.New(util.GetConfig("database", "url"))
	if err != nil {
		panic(err)
	}

	if _, err = db.Signin(map[string]interface{}{
		"user": util.GetConfig("database", "user"),
		"pass": util.GetConfig("database", "pass"),
	}); err != nil {
		panic(err)
	}

}

func Close() {
	db.Close()

}
