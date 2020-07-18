package model

const (
	AuthorizeCodeUri OauthUriType = iota
	AccessTokenUri
	RefreshTokenUri
	OpenIdUri
	UserInfoUri
)

type OauthUriType int

type Oauth struct {
	ClientId         string `json:"appid"`              // appid
	ClientSecret     string `json:"appsecret"`          // AppSecret
	CallbackUri      string `json:"callback_uri"`       // 服务器回调地址
	AuthorizeCodeUri string `json:"authorize_code_uri"` // 请求code地址
	AccessTokenUri   string `json:"access_token_uri"`   // 请求access_token地址
	RefreshTokenUri  string `json:"refresh_token_uri"`  // 请求refresh_token地址
	OpenIdUri        string `json:"openid_uri"`         // 请求openid地址
	UserInfoUri      string `json:"user_info_uri"`      // 请求用户信息地址
}

type OauthToken struct {
	AccessToken  string `json:"access_token"`  // 接口调用凭证 有效期（目前为 2 个小时）
	RefreshToken string `json:"refresh_token"` // 用户刷新 access_token 有效期（30 天）
	OpenId       string `json:"openid"`        // 授权用户唯一标识
	UnionId      string `json:"unionid"`       // 当且仅当该移动应用已获得该用户的 userinfo 授权时，才会出现该字段
	ExpiresIn    int    `json:"expires_in"`    // 过期时间
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

type OauthUser struct {
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
	Gender   string `json:"gender"`   // 性别 1 男 2 女
	Birthday string `json:"birthday"` // 生日
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市

	Token *OauthToken `json:"token"` // 令牌
}
