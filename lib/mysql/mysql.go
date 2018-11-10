package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	client *gorm.DB
)

func NewClient(ip string, port int32, userName, password, dbName string) (*gorm.DB, error) {
	if client != nil {
		return client, nil
	}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		userName, password, ip, port, dbName)
	//fmt.Println(url)
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
