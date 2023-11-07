package server

import (
	"context"
	"fmt"
	"net"

	"URLShorter/config"
	"URLShorter/internal/product"
	grpcservice "URLShorter/internal/product/delivery/grpc"
	"URLShorter/internal/product/repository/pg_repository"
	"URLShorter/internal/product/repository/redis_repository"
	"URLShorter/internal/product/usecase"
	pb "URLShorter/pkg/grpc"

	"google.golang.org/grpc"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Server struct {
	ctx    context.Context
	cfg    *config.Config
	logger *logrus.Logger
	pdb    *sqlx.DB
	rdb    redis.UniversalClient
	client *fiber.App
}

func NewServer(ctx context.Context, cfg *config.Config, logger *logrus.Logger, pdb *sqlx.DB, rdb redis.UniversalClient) *Server {
	return &Server{
		ctx:    ctx,
		cfg:    cfg,
		logger: logger,
		pdb:    pdb,
		rdb:    rdb,
		client: fiber.New(),
	}
}
func (s *Server) RunHttp() error {
	err := s.MapHandlers()
	if err != nil {
		return err
	}
	address := fmt.Sprintf(":%s", s.cfg.AppPort)
	s.logger.Info("Server started")
	if err = s.client.Listen(address); err != nil {
		return err
	}

	return nil
}

func (s *Server) RunGRPC() error {
	s.logger.Info("start gRPC")
	address := fmt.Sprintf(":%s", s.cfg.GrpcPort)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	var pdb product.PGRepo
	var rdb product.RedisRepo
	if s.cfg.UseOutMemoryStorage {
		pdb = pg_repository.NewPGRepo(s.pdb)
	}
	if s.cfg.UseInMemoryStorage {
		rdb = redis_repository.NewRedisRepo(s.rdb)
	}
	// Init use case
	uc := usecase.NewUseCase(s.ctx, s.logger, s.cfg, pdb, rdb)
	grpcServer := grpc.NewServer()
	grpcDelivery := grpcservice.NewDelivery(uc)
	pb.RegisterURLShortenerServer(grpcServer, grpcDelivery)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
