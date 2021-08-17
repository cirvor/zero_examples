package logic

import (
	"book/common/response/errorx"
	"book/common/response/successx"
	"book/service/user/model"
	"context"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchLogic {
	return SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req types.SearchReq) (*successx.SuccessResponse, error) {
	logx.Infof("userId: %v",l.ctx.Value("userId"))
	logx.Infof("number: %v",req.Number)

	userInfo, err := l.svcCtx.UserModelNoCache.FindOneByNumber(req.Number)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}

	return successx.NewDefaultSuccess(&types.SearchReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender.String,
	}), nil
}
