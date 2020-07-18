package goauth

import (
	"github.com/wenlincheng/goauth/model"
	"github.com/wenlincheng/goauth/wechat"
)

type GoAuth struct {
}

// 统一接口
type IOauth interface {
	// 设置链接
	SetUri(uriType model.OauthUriType, uri string)
	// 获取授权链接
	GetAuthorizeUrl(args ...string) string
	// 获取令牌
	GetAccessToken(code string) (*model.OauthToken, error)
	// 刷新令牌
	RefreshAccessToken(refreshToken string) (*model.OauthToken, error)
	// 获取用户信息
	GetUserInfo(accessToken, openId string) (*model.OauthUser, error)
}

// 微信
func (g *GoAuth) NewWechat(appid, appSecret string) *wechat.WeChat {
	weChat := wechat.WeChat{
		AppID:     appid,
		AppSecret: appSecret,
	}

	return &weChat
}

// QQ
func (g *GoAuth) NewQQ() {

}

// 微博
func (g *GoAuth) NewWeibo() {

}

// Github
func (g *GoAuth) NewGithub() {

}
