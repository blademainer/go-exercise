package gorm

import (
	"time"

	"github.com/blademainer/commons/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDb() (db *gorm.DB, err error) {
	if db, err = gorm.Open("sqlite3", "tmp/a.tmp"); err != nil {
		logger.Errorf("Failed to init dbConn! dialect: %s url: %s error: %s", "sqlite3", "tmp/gorm.dbConn", err.Error())
		return
	} else {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		db.DB().SetMaxIdleConns(100)

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		db.DB().SetMaxOpenConns(100)

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		// db.DB().SetConnMaxLifetime(time.Hour)
		db.DB().SetConnMaxLifetime(time.Duration(10) * time.Second)

		// print model logs
		db.LogMode(true)

		// rewrite log
		logger.Infof("Succeed connect db: %v", db)
		return
	}
}
