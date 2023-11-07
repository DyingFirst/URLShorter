package usecase

import (
	"URLShorter/config"
	"URLShorter/internal/product"
	"URLShorter/pkg/shorter"
	"URLShorter/pkg/validate"
	"context"
	"github.com/sirupsen/logrus"
)

type UseCase struct {
	ctx    context.Context
	logger *logrus.Logger
	cfg    *config.Config
	pdb    product.PGRepo
	rdb    product.RedisRepo
}

func NewUseCase(ctx context.Context, logger *logrus.Logger, cfg *config.Config, pdb product.PGRepo, rdb product.RedisRepo) product.UseCase {
	return &UseCase{ctx: ctx, logger: logger, cfg: cfg, pdb: pdb, rdb: rdb}
}

func (uc *UseCase) NewShort(OriginalURL string) (ShortedURL string, err error) {
	if ok := validate.ValidateURL(OriginalURL); ok {
		ShortedURL = shorter.URLToID(OriginalURL)
		uc.logger.Debugf("Url %v ", ShortedURL)
		if uc.cfg.UseInMemoryStorage {
			if err = uc.rdb.NewValue(uc.ctx, ShortedURL, OriginalURL); err != nil {
				uc.logger.Error("Can't write url to Redis ", err)
				return "", err
			}
		}
		if uc.cfg.UseOutMemoryStorage {
			if err = uc.pdb.NewValue(uc.ctx, ShortedURL, OriginalURL); err != nil {
				uc.logger.Error("Can't write url to Postgres ", err)
				return "", err
			}
		}
	} else {
		return "", err
	}
	return ShortedURL, err
}

func (uc *UseCase) GetOriginalURL(ShortedURL string) (OriginalURL string, err error) {
	uc.logger.Debugf("Shorted Url: %v ", ShortedURL)
	if uc.cfg.UseInMemoryStorage {
		OriginalURL, err = uc.rdb.GetValue(uc.ctx, ShortedURL)
		if err != nil {
			uc.logger.Error("Can't get url from Redis ", err)
		}
	}
	if uc.cfg.UseOutMemoryStorage {
		if OriginalURL == "" {
			OriginalURL, err = uc.pdb.GetValue(uc.ctx, ShortedURL)
			if err != nil {
				uc.logger.Error("Can't get url from Postgres ", err)
				return "", err
			}
		}
	}
	return OriginalURL, err
}
