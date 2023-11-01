package connections

import (
	"errors"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var lock = &sync.Mutex{}

type DB struct {
	db *gorm.DB
}

var singleInstance *gorm.DB

func GetInstance() (*gorm.DB, error) {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			dsn := os.Getenv("MYSQL_DSN")
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			//if err == nil {
			//	// sqlDB, err := db.DB()
			//	// if err != nil {
			//	// 	return nil, err
			//	// }
			//
			//	// // Check max connection mysql allowed: select @@max_connections
			//	// // Default: select @@max_connections = 151
			//	// // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
			//	// sqlDB.SetMaxIdleConns(10)
			//
			//	// // SetMaxOpenConns sets the maximum number of open connections to the database.
			//	// sqlDB.SetMaxOpenConns(100)
			//
			//	return db, nil
			//}
			//
			if err != nil {
				return nil, errors.New("error connect database")
			}
			singleInstance = db
		}
	}

	return singleInstance, nil
}
