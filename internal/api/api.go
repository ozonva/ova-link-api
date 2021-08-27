package api

import (
	"context"
	"os"

	grpczerolog "github.com/jwreagor/grpc-zerolog"

	"google.golang.org/grpc/grpclog"

	"github.com/rs/zerolog"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpc "github.com/ozonva/ova-link-api/pkg/ova-link-api"
)

type LinkApi struct {
	grpc.LinkAPIServer
}

func NewLinkAPI() grpc.LinkAPIServer {
	return &LinkApi{}
}

func (api *LinkApi) CreateLink(ctx context.Context, req *grpc.CreateLinkRequest) (*grpc.CreateLinkResponse, error) {
	grpclog.SetLoggerV2(grpczerolog.New(zerolog.New(os.Stdout)))
	grpclog.Info(req)
	res := &grpc.CreateLinkResponse{Id: 1}
	grpclog.Info(res)
	return res, nil
}

func (api *LinkApi) DescribeLink(ctx context.Context, req *grpc.DescribeLinkRequest) (*grpc.DescribeLinkResponse, error) {
	grpclog.SetLoggerV2(grpczerolog.New(zerolog.New(os.Stdout)))
	grpclog.Info(req)

	res := &grpc.DescribeLinkResponse{
		Id:          1,
		UserId:      2,
		Description: "Test",
		Url:         "https://url.com",
		Tags:        map[string]*emptypb.Empty{"tag1": {}, "tag2": {}},
		DateCreated: timestamppb.Now(),
	}
	grpclog.Info(res)
	return res, nil
}
func (api *LinkApi) ListLink(ctx context.Context, req *grpc.ListLinkRequest) (*grpc.ListLinkResponse, error) {
	grpclog.SetLoggerV2(grpczerolog.New(zerolog.New(os.Stdout)))
	grpclog.Info(req)

	item1 := &grpc.DescribeLinkResponse{
		Id:          uint64(1),
		UserId:      uint64(2),
		Description: "Test",
		Url:         "https://url.com",
		Tags:        map[string]*emptypb.Empty{"tag1": {}, "tag2": {}},
		DateCreated: timestamppb.Now(),
	}
	item2 := &grpc.DescribeLinkResponse{
		Id:          uint64(2),
		UserId:      uint64(3),
		Description: "Test 2",
		Url:         "https://url.com 2",
		Tags:        map[string]*emptypb.Empty{"tag3": {}, "tag4": {}},
		DateCreated: timestamppb.Now(),
	}

	items := make([]*grpc.DescribeLinkResponse, 0, 2)
	items = append(items, item1, item2)
	res := &grpc.ListLinkResponse{
		Items: items,
	}
	grpclog.Info(res)
	return res, nil
}
func (api *LinkApi) DeleteLink(ctx context.Context, req *grpc.DeleteLinkRequest) (*emptypb.Empty, error) {
	grpclog.SetLoggerV2(grpczerolog.New(zerolog.New(os.Stdout)))
	grpclog.Info(req)
	res := &emptypb.Empty{}
	grpclog.Info(res)

	return res, nil
}
