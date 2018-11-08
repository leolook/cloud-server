package service

import (
	"cloud-server/app/tool/service/impl"
	"cloud-server/common/pb"
)

type DBService interface {
	Create(*pb.Context, *pb.CreateOrModifyReq) (*pb.CreateOrModifyRsp, error)
	Update(*pb.Context, *pb.CreateOrModifyReq) (*pb.CreateOrModifyRsp, error)
}

var DB DBService = &impl.DB{}
