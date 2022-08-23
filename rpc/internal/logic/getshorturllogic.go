package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/hash"
	"shorturl-service/model"

	"shorturl-service/rpc/internal/svc"
	"shorturl-service/rpc/transform"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetShortUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetShortUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetShortUrlLogic {
	return &GetShortUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetShortUrlLogic) GetShortUrl(in *transform.GetShortUrlRequest) (*transform.GetShortUrlResponse, error) {
	key := hash.Md5Hex([]byte(in.Url))[:6]
	_, err := l.svcCtx.Model.Insert(l.ctx, &model.Shorturl{
		ShortUrl: key,
		Url:      in.Url,
	})
	if err != nil {
		return nil, err
	}
	return &transform.GetShortUrlResponse{
		ShortUrl: key,
	}, nil

}
