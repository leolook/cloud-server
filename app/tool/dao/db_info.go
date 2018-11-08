package dao

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

type DbInfoModeler interface {
	Save(model *DbInfoModel) error
	FindByName(name string) ([]*DbInfoModel, error)
}

type DbInfoDao struct {
	Base
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
