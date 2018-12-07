package impl

import (
	"cloud-server/app/tool/dao"
	"cloud-server/common/pb"
	"cloud-server/common/util"
	"cloud-server/lib/log"
	"cloud-server/lib/mysql"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/jinzhu/gorm"
)

type DB struct{}

var connectMp map[string]*gorm.DB

//创建数据库
func (i *DB) Create(ctx *pb.Context, req *pb.DbCreateOrModifyReq) (rsp *pb.DbCreateOrModifyRsp, err error) {

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
func (i *DB) Update(ctx *pb.Context, req *pb.DbCreateOrModifyReq) (rsp *pb.DbCreateOrModifyRsp, err error) {

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

	var field, order string
	if req.Sorter != nil && req.Sorter.Field != "" && req.Sorter.Order != "" {
		field = req.Sorter.Field
		switch req.Sorter.Order {
		case "ascend":
			{
				order = "asc"
			}
		case "descend":
			{
				order = "desc"
			}
		}
	} else {
		field = "id"
		order = "desc"
	}

	render := dao.Render()

	//获取分页数据
	basePage := &dao.BasePage{
		Start: start,
		End:   end,
		Field: field,
		Order: order,
	}
	models, err := render.DbInfo.Page(basePage)
	if err != nil {
		log.Errorf("failed to find db through page,[start=%d] [end=%d] [err=%v]", start, end, err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	//获取分页总数据
	count, err := render.DbInfo.PageCount()
	if err != nil {
		log.Errorf("failed to find total count,[err=%v]", err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	rsp = &pb.DbPageRsp{
		Page: &pb.PageRsp{
			Current:  req.PageNo,
			PageSize: req.PageSize,
			Total:    count,
		},
	}

	if models == nil || len(models) <= 0 {
		return rsp, nil
	}

	//构造数据
	tmp := make([]*pb.DbInfo, len(models))
	for i, v := range models {
		tmp[i] = v.ToDbInfo()
	}
	rsp.List = tmp

	return rsp, nil
}

//所有数据库
func (i *DB) AllName(ctx *pb.Context, req *pb.DbAllNameReq) (rsp *pb.DbAllNameRsp, err error) {

	render := dao.Render()

	//获取数据库
	models, err := render.DbInfo.AllName()
	if err != nil {
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	rsp = &pb.DbAllNameRsp{
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

//连接
func (i *DB) Connect(ctx *pb.Context, req *pb.DbConnectReq) (rsp *pb.DbConnectRsp, err error) {

	id, dbName := req.Id, req.DbName
	render := dao.Render()

	//获取数据库信息
	dbs, err := render.DbInfo.FindByID(id)
	if err != nil {
		log.Errorf("failed to find db through id,[id=%+v] [err=%v]", id, err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}

	if dbs == nil || len(dbs) <= 0 {
		log.Warnf("not find db through id,[id=%v]", id)
		return nil, pb.ToError(pb.E_NO_DATA, pb.P_NO_DATA)
	}

	ip, port, userName, password := dbs[0].Ip, dbs[0].Port, dbs[0].UserName, dbs[0].Password
	key := fmt.Sprintf("%s:%s:%s:%s:%d", ip, port, userName, dbName, id)
	key = base64.StdEncoding.EncodeToString([]byte(key))

	if connectMp == nil {
		connectMp = make(map[string]*gorm.DB)
	}

	rsp = &pb.DbConnectRsp{
		Key:    key,
		Id:     id,
		DbName: dbName,
	}

	var client *gorm.DB

	if v, ok := connectMp[key]; ok && v != nil {
		client = v
	} else {
		client, err = mysql.NewClient(ip, port, userName, password, dbName)
		if err != nil {
			log.Errorf("failed to connect db,[err=%v]", err)
			return nil, pb.ToError(pb.E_DB_CONNECT_FAIL, err.Error())
		}
		connectMp[key] = client
	}

	row, err := client.DB().Query("show tables")
	if err != nil {
		return nil, pb.ToError(pb.E_DB_CONNECT_FAIL, err.Error())
	}
	defer row.Close()
	tables := make([]string, 0)

	for row.Next() {
		var tmp string
		err = row.Scan(&tmp)
		if err != nil {
			return nil, pb.ToError(pb.E_DB_CONNECT_FAIL, err.Error())
		}
		tables = append(tables, tmp)
	}

	rsp.TableName = tables

	return rsp, nil
}

//表模型
func (i *DB) TableModel(ctx *pb.Context, req *pb.DbTableModelReq) (rsp *pb.DbTableModelRsp, err error) {

	if connectMp == nil {
		return nil, pb.ToError(pb.E_NOt_FOUND, pb.P_PLEASE_CONNECT_DB)
	}

	client, ok := connectMp[req.Key]
	if !ok {
		log.Warnf("not found db though key,[key=%s]", req.Key)
		return nil, pb.ToError(pb.E_NOt_FOUND, pb.P_NOT_FOUND_DB)
	}
	row, err := client.DB().Query(fmt.Sprintf("show full fields from %s", req.Name))
	if err != nil {
		return nil, pb.ToError(pb.E_DB_CONNECT_FAIL, err.Error())
	}
	defer row.Close()

	data := make([]*pb.Model, 0, 3)
	for row.Next() {
		var field, typ, comment, str sql.NullString
		err = row.Scan(&field, &typ, &str, &str, &str, &str, &str, &str, &comment)
		if err != nil {
			log.Error(err)
			return nil, pb.ToError(pb.E_SERVER_ERROR, err.Error())
		}
		tmp := &pb.Model{
			Field:   field.String,
			Type:    typ.String,
			Comment: comment.String,
		}
		data = append(data, tmp)
	}
	str := pb.ToModel(req.Name, data)

	rsp = &pb.DbTableModelRsp{
		Model: str,
	}

	return rsp, nil
}

//删除
func (DB) Del(ctx *pb.Context, req *pb.DbDelReq) (rsp *pb.DbDelRsp, err error) {
	render := dao.Render()
	//删除项
	err = render.DbInfo.Del(req.Ids...)
	if err != nil {
		log.Errorf("failed to del db option through id,[ids=%v] [err=%v]", req.Ids, err)
		return nil, pb.ToError(pb.E_SERVER_ERROR, pb.P_SERVER_ERROR)
	}
	rsp = &pb.DbDelRsp{}
	return rsp, nil
}
