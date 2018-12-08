package sqlite

import (
	"cloud-server/common/flag"
	"cloud-server/lib/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db  *gorm.DB
	err error
)

func InitDB() {
	addr := flag.SqliteAddr
	db, err = gorm.Open("sqlite3", addr)
	if err != nil {
		log.Fatalf("failed to open sqlite3 db,[addr=%s] [err=%v]", addr, err)
		return
	}
	log.Infof("init db:[addr=%s]", addr)
}

func Get() *gorm.DB {
	if db == nil {
		InitDB()
	}
	return db
}
