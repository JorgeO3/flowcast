// Package grpc provides the gRPC Controller for the catalog service.
package grpc

import (
	"context"

	"github.com/JorgeO3/flowcast/configs"
	pb "github.com/JorgeO3/flowcast/gen/catalog"
	uc "github.com/JorgeO3/flowcast/internal/catalog/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

// CreateAct handles the creation of a new act.
// It decodes the gRPC request body into CreateActInput, executes the create use case,
// and responds with the created act or an error.
func (c *Controller) CreateAct(ctx context.Context, in *pb.CreateActRequest) (*pb.CreateActResponse, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	act := convertPbActToEntity(in.Act)
	input := uc.CreateActInput{Act: act}

	output, err := c.CreateActUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing CreateAct use case - err", err)
		return nil, c.handleError(err)
	}

	return &pb.CreateActResponse{Id: output.ID}, nil
}

// UpdateAct handles updating an existing act.
// It decodes the gRPC request body into UpdateActInput, executes the update use case,
// and responds with the updated act or an error.
func (c *Controller) UpdateAct(ctx context.Context, in *pb.UpdateActRequest) (*pb.UpdateActResponse, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	act := convertPbActToEntity(in.Act)
	input := uc.UpdateActInput{Act: act}

	_, err := c.UpdateActUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing UpdateAct use case - err", err)
		return nil, c.handleError(err)
	}

	return &pb.UpdateActResponse{Success: true}, nil
}

// GetAct retrieves acts based on query parameters.
// - If 'id' is provided, it returns a single act.
// - If 'genre' is provided, it returns acts of that genre with pagination.
// - Otherwise, it returns all acts with pagination.
func (c *Controller) GetAct(ctx context.Context, in *pb.GetActRequest) (*pb.GetActResponse, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	rawID := in.Id
	id, err := primitive.ObjectIDFromHex(rawID)
	if err != nil {
		c.Logger.Error("Error parsing ObjectID - err", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	output, err := c.GetActByIDUC.Execute(ctx, uc.GetActByIDInput{ID: id})
	if err != nil {
		c.Logger.Error("Error executing GetActByID use case - err", err)
		return nil, c.handleError(err)
	}

	pbAct := convertEntityActToPb(output.Act)
	return &pb.GetActResponse{Act: pbAct}, nil
}

// DeleteAct handles the deletion of an act.
func (c *Controller) DeleteAct(ctx context.Context, in *pb.DeleteActRequest) (*pb.DeleteActResponse, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	rawID := in.Id
	id, err := primitive.ObjectIDFromHex(rawID)
	if err != nil {
		c.Logger.Error("Error parsing ObjectID - err", err)
		return nil, status.Error(codes.InvalidArgument, "Invalid ID")
	}

	_, err = c.DeleteActUC.Execute(ctx, uc.DeleteActInput{ID: id})
	if err != nil {
		c.Logger.Error("Error executing DeleteAct use case - err", err)
		return nil, c.handleError(err)
	}

	return &pb.DeleteActResponse{Success: true}, nil
}

// GetActs retrieves acts based on query parameters.
// - If 'genre' is provided, it returns acts of that genre with pagination.
// - Otherwise, it returns all acts with pagination.
func (c *Controller) GetActs(ctx context.Context, in *pb.GetActsRequest) (*pb.GetActsResponse, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	input := uc.GetActsInput{
		Limit:  int64(in.Limit),
		Offset: int64(in.Offset),
		Genre:  in.Genre,
	}

	output, err := c.GetActsUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing GetActs use case - err", err)
		return nil, c.handleError(err)
	}

	pbActs := convertEntityActsToPb(output.Acts)
	return &pb.GetActsResponse{Acts: pbActs}, nil
}

// CreateMany handles the creation of multiple acts.
// It decodes the gRPC request body into CreateManyInput, executes the create many use case,
// and responds with the created act ids or an error.
func (c *Controller) CreateMany(ctx context.Context, in *pb.CreateManyRequest) (*pb.CreateManyResponse, error) {
	ctx, cancel := withTimeout(ctx)
	defer cancel()

	acts := convertPbActsToEntity(in.Acts)
	input := uc.CreateManyInput{Acts: acts}

	output, err := c.CreateManyUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing CreateMany use case - err", err)
		return nil, c.handleError(err)
	}

	return &pb.CreateManyResponse{Ids: output.IDs}, nil
}
