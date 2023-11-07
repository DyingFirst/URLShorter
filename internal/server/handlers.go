package server

import (
	"URLShorter/internal/product"
	"URLShorter/internal/product/delivery/http"
	"URLShorter/internal/product/repository/pg_repository"
	"URLShorter/internal/product/repository/redis_repository"
	"URLShorter/internal/product/usecase"
)

func (s *Server) MapHandlers() error {
	// Init repos
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
	deliveryHttp := http.NewDelivery(s.logger, uc)
	//Map routes

	http.MapRoutes(s.client, deliveryHttp)
	return nil
}
