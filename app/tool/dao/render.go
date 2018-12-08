package dao

import (
	"cloud-server/lib/sqlite"
)

type Rend struct {
	DbInfo DbInfoModeler //数据库
}

var rend *Rend

func Render() *Rend {

	if rend != nil {
		return rend
	}

	base := Base{
		DB: sqlite.Get(),
	}

	rend = &Rend{
		DbInfo: &DbInfoDao{
			Base: base,
		},
	}

	return rend
}
