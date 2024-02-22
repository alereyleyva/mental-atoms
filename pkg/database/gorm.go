package database

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var connection *gorm.DB

const (
	ENV_DRIVER_VAR = "DRIVER"
	ENV_DSN_VAR    = "DSN"
	SQLITE_DRIVER  = "sqlite"
)

func GetConnectionInstance() *gorm.DB {
	if nil == connection {
		log.Fatal(ConnectDatabase())
	}

	return connection
}

func ConnectDatabase() error {
	var dialector gorm.Dialector
	driver := os.Getenv(ENV_DRIVER_VAR)
	dsn := os.Getenv(ENV_DSN_VAR)

	switch driver {
	case SQLITE_DRIVER:
		dialector = sqlite.Open(dsn)
	default:
		return errors.New(fmt.Sprintf("specified driver not found or not implemented yet: %s", driver))
	}

	db, err := gorm.Open(dialector, &gorm.Config{})

	if nil != err {
		return fmt.Errorf("error connecting to database: %s", err)
	}

	log.Printf("successfully connected to database: %s", dialector.Name())

	if nil != err {
		return fmt.Errorf("error migrating the database: %s", err)
	}

	connection = db

	return nil
}
