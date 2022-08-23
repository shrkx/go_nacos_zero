package logic

import (
	"context"
	"shorturl-service/rpc/transform"

	"shorturl-service/api/internal/svc"
	"shorturl-service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLongUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLongUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLongUrlLogic {
	return &GetLongUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLongUrlLogic) GetLongUrl(req *types.GetLongUrlRequest) (resp *types.GetLongUrlResponse, err error) {
	ret, err := l.svcCtx.Transformer.GetLongUrl(l.ctx, &transform.GetLongUrlRequest{
		ShortUrl: req.ShortUrl,
	})
	if err != nil {
		return nil, err
	}
	return &types.GetLongUrlResponse{
		Url: ret.Url,
	}, nil
}
