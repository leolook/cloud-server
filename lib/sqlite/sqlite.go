package sqlite

import (
	"cloud-server/lib/log"
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	sqliteAddr string
	Db         *gorm.DB
)

func init() {

	flag.StringVar(&sqliteAddr, "sqliteAddr", "tool.db", "sqlite address")
	flag.Parse()

	db, err := gorm.Open("sqlite3", sqliteAddr)
	if err != nil {
		log.Fatalf("failed to open sqlite db,[addr=%s] [err=%v]", sqliteAddr, err)
		return
	}
	Db = db
}
