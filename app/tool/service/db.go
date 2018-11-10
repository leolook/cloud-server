package service

import (
	"cloud-server/app/tool/service/impl"
	"cloud-server/common/pb"
)

type DBService interface {
	Create(*pb.Context, *pb.CreateOrModifyReq) (*pb.CreateOrModifyRsp, error)
	Update(*pb.Context, *pb.CreateOrModifyReq) (*pb.CreateOrModifyRsp, error)
	Detail(*pb.Context, *pb.DbDetailByIDReq) (*pb.DbDetailByIDRsp, error)
	Page(*pb.Context, *pb.DbPageReq) (*pb.DbPageRsp, error)
	AllName(*pb.Context, *pb.AllNameReq) (*pb.AllNameRsp, error)
}

var DB DBService = &impl.DB{}
