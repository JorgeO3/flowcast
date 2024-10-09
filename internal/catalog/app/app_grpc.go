//go:build grpc

// Package app provides the entry point to the catalog service.
package app

import (
	"fmt"
	"net"
	"time"

	"github.com/JorgeO3/flowcast/configs"
	pb "github.com/JorgeO3/flowcast/gen/catalog"
	c "github.com/JorgeO3/flowcast/internal/catalog/controller/grpc"
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
	uc "github.com/JorgeO3/flowcast/internal/catalog/usecase"
	grpcu "github.com/JorgeO3/flowcast/pkg/grpcutils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/mongodb"
	"github.com/JorgeO3/flowcast/pkg/validator"
	"golang.org/x/time/rate"
)

// Run starts the gRPC server for the catalog service.
func Run(cfg *configs.CatalogConfig, logg logger.Interface) {
	logg.Info("Starting catalog service")

	mg, err := mongodb.New(cfg.DatabaseURL)
	if err != nil {
		logg.Fatal("mongo connection error: %v", err)
	}
	defer mg.Close()

	logg.Debug("Connected to MongoDB")

	db := mg.Client.Database(entity.Database)
	actRepo := repository.NewMongoActRepository(db, entity.ActCollection)

	val := validator.New()

	createActUC := uc.NewCreateAct(
		uc.WithCreateActLogger(logg),
		uc.WithCreateActValidator(val),
		uc.WithCreateActRepository(actRepo),
	)

	updateActUC := uc.NewUpdateAct(
		uc.WithUpdateActLogger(logg),
		uc.WithUpdateActValidator(val),
		uc.WithUpdateActRepository(actRepo),
	)

	getActByIDUC := uc.NewGetActByID(
		uc.WithGetAcByIDLogger(logg),
		uc.WithGetAcByIDValidator(val),
		uc.WithGetAcByIDRepository(actRepo),
	)

	deleteActUC := uc.NewDeleteAct(
		uc.WithDeleteActLogger(logg),
		uc.WithDeleteActValidator(val),
		uc.WithDeleteActRepository(actRepo),
	)

	createManyUC := uc.NewCreateMany(
		uc.WithCreateManyLogger(logg),
		uc.WithCreateManyValidator(val),
		uc.WithCreateManyRepository(actRepo),
	)

	getActsUC := uc.NewGetActs(
		uc.WithGetActsLogger(logg),
		uc.WithGetActsValidator(val),
		uc.WithGetActsRepository(actRepo),
	)

	controller := c.New(
		c.WithConfig(cfg),
		c.WithLogger(logg),
		c.WithGetActsUC(getActsUC),
		c.WithCreateActUC(createActUC),
		c.WithUpdateActUC(updateActUC),
		c.WithDeleteActUC(deleteActUC),
		c.WithGetActByIDUC(getActByIDUC),
		c.WithCreateManyUC(createManyUC),
	)

	// Create a new rate limiter
	limiter := rate.NewLimiter(5, 10)

	// Create a new gRPC server using the new middleware package
	server := grpcu.NewServer(
		grpcu.WithLogger(logg),
		grpcu.WithUnaryInterceptor(grpcu.RequestIDInterceptor()),
		grpcu.WithUnaryInterceptor(grpcu.RealIPInterceptor()),
		grpcu.WithUnaryInterceptor(grpcu.RecoveryInterceptor()),
		grpcu.WithUnaryInterceptor(grpcu.RequestSizeLimitInterceptor(1024*1024)), // 1 MB
		grpcu.WithUnaryInterceptor(grpcu.TimeoutInterceptor(5*time.Second)),
		grpcu.WithUnaryInterceptor(grpcu.RateLimitInterceptor(limiter)),
		grpcu.WithUnaryInterceptor(grpcu.JWTAuthInterceptor(myValidateTokenFunc)),
		grpcu.WithUnaryInterceptor(grpcu.LoggerErrorHandlingInterceptor(logg)),
	)

	// Register the CatalogService server implementation
	pb.RegisterCatalogServiceServer(server, controller)

	// Server address
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	// Start listening on the configured port
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logg.Fatal("Failed to listen on %s: %v", addr, err)
	}

	logg.Info("gRPC server listening on " + addr)

	// Start the gRPC server
	if err := server.Serve(lis); err != nil {
		logg.Fatal("Failed to serve gRPC server: %v", err)
	}
}

// Función de validación de token de ejemplo
func myValidateTokenFunc(_ string) error {
	// Implementa la validación de tu token aquí
	return nil
}
