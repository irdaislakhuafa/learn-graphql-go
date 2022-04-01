package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DBConnection struct {
	Username string
	Password string
	Driver   string
	DBName   string
	DBPort   string
	db       *sqlx.DB
}

func (me *DBConnection) Connect() {
	var err error
	me.db, err = sqlx.Connect(me.Driver, fmt.Sprintf("%s:%s@(localhost:%s)/%s", me.Username, me.Password, me.DBPort, me.DBName))

	if err != nil {
		panic(fmt.Sprintf("Error while connecting to database \"%s\" because error : %v", me.DBName, err.Error()))
	}

	log.Printf("Success connected to database \"%s\" on port %s \n", me.DBName, me.DBPort)
}
func (me *DBConnection) GetConnection() *sqlx.DB {
	return me.db
}
