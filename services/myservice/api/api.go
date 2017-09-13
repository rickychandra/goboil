package api

import (
	"github.com/rickychandra/goboil/common/ctxutil"
	"github.com/rickychandra/goboil/common/logger"
	"github.com/rickychandra/goboil/common/pb"
	"github.com/rickychandra/goboil/common/stringutil"
	pb2 "github.com/rickychandra/goboil/shared/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ pb2.MyServiceServer = (*myServiceServer)(nil)

type myServiceServer struct{}

func NewMyServiceServer() pb2.MyServiceServer {
	return new(myServiceServer)
}

func (s *myServiceServer) Hello(ctx context.Context, req *pb2.HelloRequest) (
	*pb2.HelloResponse, error) {

	//Must preprocess context.
	ctx = ctxutil.PreprocessAPICtx(ctx)

	logger.Info(ctx, stringutil.Sprintf("hello request: %v", req))
	res := new(pb2.HelloResponse)
	res.Ok = false

	return res, nil
}

func (s *myServiceServer) HelloError(ctx context.Context, req *pb2.HelloErrorRequest) (
	*common.Empty, error) {

	//Must preprocess context.
	ctx = ctxutil.PreprocessAPICtx(ctx)

	logger.Info(ctx, stringutil.Sprintf("hello error request: %v", req))

	//Returns error.
	return nil, status.Error(codes.Unknown, "an expected error")
}
