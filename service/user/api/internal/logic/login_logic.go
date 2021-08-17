package logic

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"

	"zero_examples/common/response/errorx"
	"zero_examples/common/response/successx"
	"zero_examples/service/user/api/internal/svc"
	"zero_examples/service/user/api/internal/types"
	"zero_examples/service/user/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(req types.LoginReq) (*successx.SuccessResponse, error) {
	if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewDefaultError("参数错误")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(req.Username)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewDefaultError("用户名不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, errorx.NewDefaultError("用户密码不正确")
	}

	// ---start---
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, errorx.NewDefaultError("密钥生成失败")
	}
	// ---end---

	return successx.NewDefaultSuccess(&types.LoginReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender.String,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}), nil
}
