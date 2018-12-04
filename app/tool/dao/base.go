package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
)

type BasePage struct {
	Start, End   int32
	Field, Order string
}

type Base struct {
	DB *gorm.DB
}

func (b *Base) checkRow(row *gorm.DB) (bool, error) {

	if row.RecordNotFound() {
		return false, nil
	}

	if row.Error != nil {
		return false, row.Error
	}

	return true, nil
}

func (b *Base) modifyColumn(model interface{}) (string, []interface{}) {

	v := reflect.ValueOf(model).Elem()
	t := reflect.TypeOf(model).Elem()

	inter := make([]interface{}, 0)
	attr := make([]string, 0)

	for i := 0; i < t.NumField(); i++ {

		val := v.Field(i).Interface()
		tmp := fmt.Sprintf("%v", val)
		if tmp == "" || tmp == "0" {
			continue
		}

		str := t.Field(i).Tag.Get("gorm")
		if str == "" || !strings.Contains(str, "column:") {
			continue
		}

		name := strings.Split(str, ":")[1]
		if name == "id" {
			continue
		}

		attr = append(attr, name)
		inter = append(inter, val)
	}

	if len(attr) == 0 {
		return "", nil
	}

	var str string
	for i, v := range attr {
		str += v + "=?"
		if i == len(attr)-1 {
			break
		}
		str += ","
	}
	return str, inter
}
