package service

import (
	"cloud-server/app/tool/service/impl"
	"cloud-server/common/pb"
)

type DBService interface {
	Create(*pb.Context, *pb.CreateReq) (*pb.CreateRsp, error)
}

var DB DBService = &impl.DB{}
