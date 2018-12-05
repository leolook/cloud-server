package sqlite

import (
	"cloud-server/common/flag"
	"cloud-server/lib/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	Db *gorm.DB
)

func init() {
	db, err := gorm.Open("sqlite3", flag.SqliteAddr)
	if err != nil {
		log.Fatalf("failed to open sqlite3 db,[addr=%s] [err=%v]", flag.SqliteAddr, err)
		return
	}
	Db = db
}
