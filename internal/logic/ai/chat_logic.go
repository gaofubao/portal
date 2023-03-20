package ai

import (
	"context"

	"portal/internal/svc"
	"portal/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatLogic {
	return &ChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatLogic) Chat(req *types.ChatReq) (resp *types.ChatResp, err error) {
	resp = &types.ChatResp{}

	reply, err := l.svcCtx.Model.Reply(req.UserId, req.Query)
	if err != nil {
		return
	}

	resp.Answer = reply
	return
}
