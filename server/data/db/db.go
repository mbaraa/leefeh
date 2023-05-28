package db

import (
	"fmt"
	"salsa/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var instance *gorm.DB = nil

// GetDBConnector returns a singleton mysql connection instance to the application's DB
func GetDBConnector() *gorm.DB {
	return getDBConnector("salsa")
}

// GetTestDBConnector returns a singleton mysql connection instance to the application's test DB
func GetTestDBConnector() *gorm.DB {
	return getDBConnector("salsa").Debug()
}

// getDBConnector returns a singleton mysql connection instance
func getDBConnector(dbName string) *gorm.DB {
	if instance == nil {
		var err error
		createDBDsn := fmt.Sprintf("%s:%s@tcp(%s)/", config.DBUser(), config.DBPassword(), config.DBHost())
		database, err := gorm.Open(mysql.Open(createDBDsn), &gorm.Config{})

		_ = database.Exec("CREATE DATABASE IF NOT EXISTS " + dbName + ";")

		instance, err = gorm.Open(mysql.New(mysql.Config{
			DriverName: "mysql",
			DSN:        fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True&loc=Local", config.DBUser(), config.DBPassword(), config.DBHost(), dbName),
		}))
		if err != nil {
			panic(err)
		}
	}
	return instance
}

func InitTables() {
	if instance != nil {
		err := instance.AutoMigrate()
		if err != nil {
			panic(err)
		}
	} else {
		panic("No DB connection was initialized")
	}
}
