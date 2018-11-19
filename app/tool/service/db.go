package service

import (
	"cloud-server/app/tool/service/impl"
	"cloud-server/common/pb"
)

type DBService interface {
	Create(*pb.Context, *pb.DbCreateOrModifyReq) (*pb.DbCreateOrModifyRsp, error)
	Update(*pb.Context, *pb.DbCreateOrModifyReq) (*pb.DbCreateOrModifyRsp, error)
	Detail(*pb.Context, *pb.DbDetailByIDReq) (*pb.DbDetailByIDRsp, error)
	Page(*pb.Context, *pb.DbPageReq) (*pb.DbPageRsp, error)
	AllName(*pb.Context, *pb.DbAllNameReq) (*pb.DbAllNameRsp, error)
	Connect(*pb.Context, *pb.DbConnectReq) (*pb.DbConnectRsp, error)
	TableModel(*pb.Context, *pb.DbTableModelReq) (*pb.DbTableModelRsp, error)
}

var DB DBService = &impl.DB{}
