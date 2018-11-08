package impl

import (
	"cloud-server/app/tool/dao"
	"cloud-server/common/pb"
	"cloud-server/common/util"
	"cloud-server/lib/log"
)

type DB struct{}

//创建数据库
func (i *DB) Create(ctx *pb.Context, req *pb.CreateReq) (rsp *pb.CreateRsp, err error) {

	log.Infof("%v", *ctx)

	render := dao.Render()

	model := &dao.DbInfoModel{
		Ip:       req.Db.Ip,
		Name:     req.Db.Name,
		Port:     int(req.Db.Port),
		UserName: req.Db.UserName,
		Password: req.Db.Password,
		CreateAt: util.Time.Unix(),
		UpdateAt: util.Time.Unix(),
	}

	err = render.DbInfo.Save(model)
	if err != nil {
		return nil, pb.ToError(pb.E_SERVER_ERROR, err.Error())
	}

	return nil, nil
}
