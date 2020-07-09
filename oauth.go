package goauth

type GoAuth struct {
}

// 统一接口
type IOauth interface {
	SetUri(uriType OauthUriType, uri string)

	GetAuthorizeUrl(args ...string) string
	// 获取令牌
	GetAccessToken(code string) (*OauthToken, error)
	// 刷新令牌
	RefreshAccessToken(refreshToken string) (*OauthToken, error)
	GetUserInfo(accessToken, openId string) (*OauthUser, error)
}

// 微信
func (g *GoAuth) NewWeChat() {

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
