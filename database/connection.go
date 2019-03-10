package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func InitDatabase() (db *gorm.DB, err error) {
	var connectionString string
	dbDriver := "mysql"
	connectionString = buildMySqlConnectionString()
	logrus.Infoln("Connection "+connectionString)
	db, err = openConnection(dbDriver, connectionString)
	return
}

func openConnection(dbDriver, connection string) (db *gorm.DB, err error) {
	db, err = gorm.Open(dbDriver, connection)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func buildMySqlConnectionString() (connectionString string) {
	connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"root",
		"tax_calculator")

	return
}