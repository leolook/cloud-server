package impl

import (
	"cloud-server/app/tool/dao"
	"cloud-server/common/pb"
	"cloud-server/common/util"
	"cloud-server/lib/log"
)

type DB struct{}

//创建数据库
func (i *DB) Create(ctx *pb.Context, req *pb.CreateOrModifyReq) (rsp *pb.CreateOrModifyRsp, err error) {

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

//修改数据库
func (i *DB) Update(ctx *pb.Context, req *pb.CreateOrModifyReq) (rsp *pb.CreateOrModifyRsp, err error) {

	db := req.Db

	render := dao.Render()

	//获取数据库信息
	dbs, err := render.DbInfo.FindByID(db.Id)
	if err != nil {
		log.Errorf("failed to find by id,[id=%v] [err=%v]", db.Id, err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	if dbs == nil {
		log.Errorf("not find db info through id,[id=%d]", db.Id)
		return nil, pb.ToError(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ID)
	}
	tmp := dbs[0]

	model := new(dao.DbInfoModel)
	if db.Name != tmp.Name {
		model.Name = db.Name
	}
	if db.Ip != tmp.Ip {
		model.Ip = db.Ip
	}
	if db.Port != tmp.Port {
		model.Port = db.Port
	}
	if db.UserName != tmp.UserName {
		model.UserName = db.UserName
	}
	if db.Password != tmp.Password {
		model.Password = db.Password
	}
	model.UpdateAt = util.Time.Unix()
	model.ID = db.Id

	//校验是否有重复名
	if model.Name != "" {
		dbs, err = render.DbInfo.FindByName(db.Name)
		if err != nil {
			log.Errorf("failed to find by name,[name=%s] [err=%v]", db.Name, err)
			return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
		}

		if dbs != nil || len(dbs) > 0 {
			return nil, pb.ToError(pb.E_DB_NAME_REPEATED, pb.P_NAME_REPEATED)
		}
	}

	//修改数据
	err = render.DbInfo.Modify(model)
	if err != nil {
		log.Errorf("failed to modify db info,[err=%v]", err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}
	return nil, nil
}

//数据库详情
func (i *DB) Detail(ctx *pb.Context, req *pb.DbDetailByIDReq) (rsp *pb.DbDetailByIDRsp, err error) {

	id := req.Id

	render := dao.Render()

	//获取数据库信息
	dbs, err := render.DbInfo.FindByID(id)
	if err != nil {
		log.Errorf("failed to find db info through id,[id=%+v] [err=%v]", id, err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	if dbs == nil || len(dbs) <= 0 {
		log.Warnf("not find db info through id,[id=%+v]", id)
		return nil, pb.ToError(pb.E_INVALID_ARGUMENT, pb.P_INVALID_ID)
	}

	db := dbs[0]

	//构造返回数据
	rsp = &pb.DbDetailByIDRsp{
		Db: &pb.DbInfo{
			Name:     db.Name,
			Ip:       db.Ip,
			Port:     db.Port,
			UserName: db.UserName,
			Password: db.Password,
			Id:       db.ID,
		},
	}

	return rsp, nil
}

//数据库分页
func (i *DB) Page(ctx *pb.Context, req *pb.DbPageReq) (rsp *pb.DbPageRsp, err error) {

	var start, end int32
	start = (req.PageNo - 1) * req.PageSize
	end = req.PageSize

	render := dao.Render()

	//获取分页数据
	models, err := render.DbInfo.Page(start, end)
	if err != nil {
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	rsp = &pb.DbPageRsp{
		IsMore:   1, //1-否 2-是
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
	}

	if models == nil || len(models) <= 0 {
		return rsp, nil
	}

	//构造数据
	tmp := make([]*pb.DbInfo, len(models))
	for i, v := range models {
		tmp[i] = v.ToDbInfo()
	}
	rsp.Dbs = tmp
	rsp.IsMore = 2

	return rsp, nil
}

//所有数据库
func (i *DB) AllName(ctx *pb.Context, req *pb.AllNameReq) (rsp *pb.AllNameRsp, err error) {

	render := dao.Render()

	//获取数据库
	models, err := render.DbInfo.AllName()
	if err != nil {
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	rsp = &pb.AllNameRsp{
		Dbs: nil,
	}

	if models == nil || len(models) <= 0 {
		return rsp, nil
	}

	//构造数据
	tmp := make([]*pb.DbInfo, len(models))
	for i, v := range models {
		tmp[i] = v.ToDbInfo()
	}
	rsp.Dbs = tmp

	return rsp, nil
}
