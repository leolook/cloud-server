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

	render := dao.Render()

	//校验是否有重复名
	dbs, err := render.DbInfo.FindByName(req.Db.Name)
	if err != nil {
		log.Errorf("failed to find by name,[name=%s] [err=%v]", req.Db.Name, err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	if dbs != nil || len(dbs) > 0 {
		return nil, pb.ToError(pb.E_DB_NAME_REPEATED, pb.P_NAME_REPEATED)
	}

	currentUnix := util.Time.Unix()
	model := &dao.DbInfoModel{
		Ip:       req.Db.Ip,
		Name:     req.Db.Name,
		Port:     req.Db.Port,
		UserName: req.Db.UserName,
		Password: req.Db.Password,
		CreateAt: currentUnix,
		UpdateAt: currentUnix,
	}

	//保存数据
	err = render.DbInfo.Save(model)
	if err != nil {
		log.Errorf("failed to save db info,[err=%v]", err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	return nil, nil
}
