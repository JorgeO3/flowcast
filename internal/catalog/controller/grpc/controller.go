// Package grpc provides the gRPC Controller for the catalog service.
package grpc

import (
	"context"

	"github.com/JorgeO3/flowcast/configs"
	pb "github.com/JorgeO3/flowcast/gen/catalog"
	uc "github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
)

// Controller handles gRPC requests for the catalog service and delegates operations to use cases.
type Controller struct {
	pb.UnimplementedCatalogServiceServer

	GetActsUC    *uc.GetActsUC
	DeleteActUC  *uc.DeleteActUC
	CreateActUC  *uc.CreateActUC
	UpdateActUC  *uc.UpdateActUC
	GetActByIDUC *uc.GetActByIDUC
	CreateManyUC *uc.CreateManyUC

	Logger logger.Interface
	Cfg    *configs.CatalogConfig
}

func (c *Controller) CreateAct(ctx context.Context, in *pb.CreateActRequest) (*pb.CreateActResponse, error) {
	return nil, nil
}

func (c *Controller) CreateMany(ctx context.Context, in *pb.CreateManyRequest) (*pb.CreateManyResponse, error) {
	return nil, nil
}

func (c *Controller) DeleteAct(ctx context.Context, in *pb.DeleteActRequest) (*pb.DeleteActResponse, error) {
	return nil, nil
}

func (c *Controller) GetAct(ctx context.Context, in *pb.GetActRequest) (*pb.GetActResponse, error) {
	return nil, nil
}

func (c *Controller) GetActs(ctx context.Context, in *pb.GetActsRequest) (*pb.GetActsResponse, error) {
	return nil, nil
}

func (c *Controller) UpdateAct(ctx context.Context, in *pb.UpdateActRequest) (*pb.UpdateActResponse, error) {
	return nil, nil
}
