package dao

type DbInfoModel struct {
	ID       int64  `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Ip       string `gorm:"column:ip"`
	Port     int    `gorm:"column:port"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	CreateAt int64  `gorm:"column:create_at"`
	UpdateAt int64  `gorm:"column:update_at"`
	DeleteAt int64  `gorm:"column:delete_at"`
}

type DbInfoModeler interface {
	Save(model *DbInfoModel) error
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
