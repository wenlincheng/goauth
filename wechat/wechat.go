package wechat

/**

微信


*/

const (
	// 获取授权凭证
	AccessTokenUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code"
	// 刷新/续期授权凭证
	RefreshTokenUrl = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=APPID&grant_type=refresh_token&refresh_token=REFRESH_TOKEN"
	// 检验授权凭证
	InvalidCheckUrl = "https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID"
	// 获取用户个人信息
	UserInfoUrl = "https://api.weixin.qq.com/sns/userinfo?access_token=ACCESS_TOKEN&openid=OPENID"
)

type WeChat struct {
	AppID     string
	AppSecret string
}

//
// 获取授权凭证
//
//
func (w *WeChat) GetAccessToken() *AccessTokenResp {

	return nil
}

//
// 刷新/续期授权凭证
//
//
func (w *WeChat) RefreshAccessToken() *AccessTokenResp {

	return nil
}

//
// 获取用户信息
//
//
func (w *WeChat) GetUserInfo() *UserInfoResp {

	return nil
}
