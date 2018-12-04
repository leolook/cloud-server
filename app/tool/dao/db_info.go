package dao

import (
	"cloud-server/common/pb"
	"fmt"
)

type DbInfoModel struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Ip       string `gorm:"column:ip"`
	Port     int32  `gorm:"column:port"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	CreateAt int64  `gorm:"column:create_at"`
	UpdateAt int64  `gorm:"column:update_at"`
	DeleteAt int64  `gorm:"column:delete_at"`
}

func (t *DbInfoModel) ToDbInfo() *pb.DbInfo {
	return &pb.DbInfo{
		Name:     t.Name,
		Ip:       t.Ip,
		Port:     t.Port,
		UserName: t.UserName,
		Password: t.Password,
		Id:       t.ID,
	}
}

type DbInfoModeler interface {
	Save(model *DbInfoModel) error
	FindByName(name string) ([]*DbInfoModel, error)
	FindByID(id ...int64) ([]*DbInfoModel, error)
	Modify(model *DbInfoModel) error
	Page(v *BasePage) ([]*DbInfoModel, error)
	PageCount() (int32, error)
	AllName() ([]*DbInfoModel, error)
	Del(ids ...int64) error
}

type DbInfoDao struct {
	Base
}

//删除
func (t *DbInfoDao) Del(ids ...int64) error {

	str := "update db_info set delete_at=1 where "
	if len(ids) > 1 {
		str += "id in(?)"
	} else {
		str += "id=?"
	}

	if err := t.DB.Exec(str, ids).Error; err != nil {
		return err
	}
	return nil
}

func (t *DbInfoDao) AllName() ([]*DbInfoModel, error) {
	str := "select id,name from db_info where delete_at=0 order by id desc "

	var models []*DbInfoModel
	row := t.DB.Raw(str).Scan(&models)

	ok, err := t.checkRow(row)
	if ok && models != nil && len(models) > 0 {
		return models, nil
	}
	return nil, err
}

func (t *DbInfoDao) Page(v *BasePage) ([]*DbInfoModel, error) {

	str := fmt.Sprintf("select * from db_info where delete_at=0 order by %s %s limit ?,?", v.Field, v.Order)

	var models []*DbInfoModel
	row := t.DB.Raw(str, v.Start, v.End).Scan(&models)

	ok, err := t.checkRow(row)
	if ok && models != nil && len(models) > 0 {
		return models, nil
	}
	return nil, err
}

func (t *DbInfoDao) PageCount() (int32, error) {

	str := "select count(id) from db_info where delete_at=0"

	var count int32
	row := t.DB.Raw(str).Count(&count)

	ok, err := t.checkRow(row)
	if ok && count > 0 {
		return count, nil
	}
	return 0, err
}

func (t *DbInfoDao) Modify(model *DbInfoModel) error {

	str, val := t.modifyColumn(model)
	if str == "" {
		return nil
	}
	str = fmt.Sprintf("update db_info set %s where id=?", str)
	val = append(val, model.ID)

	row := t.DB.Exec(str, val...)
	if row.Error != nil {
		return row.Error
	}
	return nil
}

func (t *DbInfoDao) FindByID(id ...int64) ([]*DbInfoModel, error) {

	str := "select id,name,ip,port,user_name,password from db_info where delete_at=0 and "
	if len(id) > 1 {
		str += "id in(?)"
	} else {
		str += "id =?"
	}

	var models []*DbInfoModel
	row := t.DB.Raw(str, id).Scan(&models)

	ok, err := t.checkRow(row)
	if ok && models != nil && len(models) > 0 {
		return models, nil
	}
	return nil, err
}

func (t *DbInfoDao) Save(model *DbInfoModel) error {
	row := t.DB.Table("db_info").Save(model)
	if row.Error != nil {
		return row.Error
	}
	return nil
}

func (t *DbInfoDao) FindByName(name string) ([]*DbInfoModel, error) {

	str := "select * from db_info where name =?"
	var models []*DbInfoModel
	row := t.DB.Raw(str, name).Scan(&models)

	ok, err := t.checkRow(row)
	if ok && models != nil && len(models) > 0 {
		return models, nil
	}
	return nil, err
}
