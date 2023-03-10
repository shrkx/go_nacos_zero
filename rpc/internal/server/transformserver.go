// Code generated by goctl. DO NOT EDIT!
// Source: transform.proto

package server

import (
	"context"

	"shorturl-service/rpc/internal/logic"
	"shorturl-service/rpc/internal/svc"
	"shorturl-service/rpc/transform"
)

type TransformServer struct {
	svcCtx *svc.ServiceContext
	transform.UnimplementedTransformServer
}

func NewTransformServer(svcCtx *svc.ServiceContext) *TransformServer {
	return &TransformServer{
		svcCtx: svcCtx,
	}
}

func (s *TransformServer) GetShortUrl(ctx context.Context, in *transform.GetShortUrlRequest) (*transform.GetShortUrlResponse, error) {
	l := logic.NewGetShortUrlLogic(ctx, s.svcCtx)
	return l.GetShortUrl(in)
}

func (s *TransformServer) GetLongUrl(ctx context.Context, in *transform.GetLongUrlRequest) (*transform.GetLongUrlResponse, error) {
	l := logic.NewGetLongUrlLogic(ctx, s.svcCtx)
	return l.GetLongUrl(in)
}
