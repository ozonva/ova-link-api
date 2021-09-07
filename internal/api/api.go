package api

import (
	"context"

	"github.com/ozonva/ova-link-api/internal/metrics"

	"github.com/ozonva/ova-link-api/internal/utils"

	"github.com/ozonva/ova-link-api/internal/kafka"

	"github.com/ozonva/ova-link-api/internal/repo"

	"github.com/ozonva/ova-link-api/internal/link"

	grpczerolog "github.com/jwreagor/grpc-zerolog"
	"github.com/ozonva/ova-link-api/internal/flusher"
	"github.com/ozonva/ova-link-api/internal/saver"

	"google.golang.org/grpc/grpclog"

	"github.com/rs/zerolog"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	grpc "github.com/ozonva/ova-link-api/pkg/ova-link-api"
)

type LinkAPI struct {
	grpc.LinkAPIServer
	repo     repo.Repo
	saver    saver.Saver
	logger   zerolog.Logger
	producer kafka.Producer
	metrics  metrics.Metrics
}

func NewLinkAPI(repo repo.Repo, logger zerolog.Logger, producer kafka.Producer, metrics metrics.Metrics) grpc.LinkAPIServer {
	api := &LinkAPI{}
	api.repo = repo
	api.saver = saver.NewTimeOutSaver(10, flusher.NewFlusher(3, api.repo), 1)
	api.logger = logger
	api.producer = producer
	api.metrics = metrics

	return api
}

func (api *LinkAPI) CreateLink(ctx context.Context, req *grpc.CreateLinkRequest) (*emptypb.Empty, error) {
	grpclog.SetLoggerV2(grpczerolog.New(api.logger))
	grpclog.Info(req)

	entity := link.New(req.UserId, req.Url)
	entity.Description = req.Description
	entity.SetTagsAsSlice(req.Tags)
	api.saver.Save(*entity)

	err := api.producer.Send(kafka.Message{
		EventType: kafka.Create,
		Value:     *entity,
	})

	if err != nil {
		return nil, err
	}

	api.metrics.CreateSuccessResponseCounter()
	res := &emptypb.Empty{}
	grpclog.Info(res)
	return res, nil
}

func (api *LinkAPI) DescribeLink(ctx context.Context, req *grpc.DescribeLinkRequest) (*grpc.DescribeLinkResponse, error) {
	grpclog.SetLoggerV2(grpczerolog.New(api.logger))
	grpclog.Info(req)

	res := &grpc.DescribeLinkResponse{}
	result, err := api.repo.DescribeEntity(req.GetId())
	if err != nil {
		return res, err
	}

	res.Id = result.ID
	res.UserId = result.UserID
	res.Description = result.Description
	res.Url = result.Url
	res.Tags = result.GetTagsAsSlice()
	res.DateCreated = timestamppb.New(result.CreatedAt)

	api.metrics.DescribeSuccessResponseCounter()
	grpclog.Info(res)
	return res, nil
}

func (api *LinkAPI) ListLink(ctx context.Context, req *grpc.ListLinkRequest) (*grpc.ListLinkResponse, error) {
	grpclog.SetLoggerV2(grpczerolog.New(api.logger))
	grpclog.Info(req)

	res := &grpc.ListLinkResponse{}
	result, err := api.repo.ListEntities(*req.Limit, *req.Offset)
	if err != nil {
		return res, err
	}

	for _, entity := range result {
		resEntity := &grpc.DescribeLinkResponse{
			Id:          entity.ID,
			UserId:      entity.UserID,
			Description: entity.Description,
			Url:         entity.Url,
			Tags:        entity.GetTagsAsSlice(),
			DateCreated: timestamppb.New(entity.CreatedAt),
		}

		res.Items = append(res.Items, resEntity)
	}

	api.metrics.ListSuccessResponseCounter()
	grpclog.Info(res)
	return res, nil
}
func (api *LinkAPI) DeleteLink(ctx context.Context, req *grpc.DeleteLinkRequest) (*emptypb.Empty, error) {
	grpclog.SetLoggerV2(grpczerolog.New(api.logger))
	grpclog.Info(req)

	res := &emptypb.Empty{}
	err := api.repo.DeleteEntity(req.GetId())
	if err != nil {
		return res, err
	}

	err = api.producer.Send(kafka.Message{
		EventType: kafka.Remove,
		Value:     req.GetId(),
	})
	if err != nil {
		return nil, err
	}
	api.metrics.RemoveSuccessResponseCounter()
	grpclog.Info(res)
	return res, nil
}

func (api *LinkAPI) UpdateLink(ctx context.Context, req *grpc.UpdateLinkRequest) (*emptypb.Empty, error) {
	grpclog.SetLoggerV2(grpczerolog.New(api.logger))
	grpclog.Info(req)

	res := &emptypb.Empty{}
	entity := link.New(req.UserId, req.Url)
	entity.Description = req.Description
	entity.SetTagsAsSlice(req.Tags)

	err := api.repo.UpdateEntity(*entity, req.Id)
	if err != nil {
		return res, err
	}

	err = api.producer.Send(kafka.Message{
		EventType: kafka.Update,
		Value:     *entity,
	})
	if err != nil {
		return nil, err
	}

	api.metrics.UpdateSuccessResponseCounter()
	grpclog.Info(res)
	return res, nil
}

func (api *LinkAPI) MultiCreateLink(ctx context.Context, req *grpc.MultiCreateLinkRequest) (*emptypb.Empty, error) {
	grpclog.SetLoggerV2(grpczerolog.New(api.logger))
	grpclog.Info(req)

	links := make([]link.Link, 0, len(req.Items))
	for _, item := range req.Items {
		entity := link.New(item.UserId, item.Url)
		entity.Description = item.Description
		entity.SetTagsAsSlice(item.Tags)
		links = append(links, *entity)
	}

	bulks := utils.SliceChunkLink(links, uint(3))

	for _, bulk := range bulks {
		err := api.repo.AddEntities(bulk)
		if err != nil {
			return nil, err
		}

		err = api.producer.Send(kafka.Message{
			EventType: kafka.MultiCreate,
			Value:     bulk,
		})
		if err != nil {
			return nil, err
		}
	}

	api.metrics.MultiCreateSuccessResponseCounter()
	res := &emptypb.Empty{}
	grpclog.Info(res)
	return res, nil
}
