package logic

import (
	"context"
	"shorturl/rpc/transform/transformer"

	"github.com/zeromicro/go-zero/core/logx"
	"shorturl/api/internal/svc"
	"shorturl/api/internal/types"
)

type ShortenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShortenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ShortenLogic {
	return &ShortenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShortenLogic) Shorten(req *types.ShortenReq) (resp *types.ShortenResp, err error) {
	rpcResp, err := l.svcCtx.Transformer.Shorten(l.ctx, &transformer.ShortenReq{
		Url: req.Url,
	})

	if err != nil {
		return nil, err
	}

	return &types.ShortenResp{
		Shorten: rpcResp.Shorten,
	}, nil
}
