#!/usr/bin/env bash
if [ ! -d pb ]
then
    mkdir -p pb
fi
if [ ! -d pb ]
then
    ln -s ../../common/pb/ proto
fi
protoc --go_out=plugins=grpc:./common/pb ./common/proto/*.proto  ./common/proto/*.proto -I./common/proto/ -I./common/proto/
cd common/pb/;
ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
