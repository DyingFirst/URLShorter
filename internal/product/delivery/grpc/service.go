package grpc

import (
	"context"

	"URLShorter/internal/product"
	pb "URLShorter/pkg/grpc"
)

type Shorter struct {
	uc product.UseCase
	pb.UnimplementedURLShortenerServer
}

func NewDelivery(uc product.UseCase) *Shorter {
	return &Shorter{uc: uc}
}

func (d *Shorter) ShortURL(_ context.Context, request *pb.ShortURLRequest) (*pb.ShortURLResponse, error) {
	shortURL, err := d.uc.NewShort(request.GetOriginURL())
	if err != nil {
		return nil, err
	}
	return &pb.ShortURLResponse{ShortURL: shortURL}, nil
}

func (d *Shorter) GetURL(_ context.Context, request *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	originURL, err := d.uc.GetOriginalURL(request.GetShortURL())
	if err != nil {
		return nil, err
	}
	return &pb.GetURLResponse{OriginURL: originURL}, nil
}
