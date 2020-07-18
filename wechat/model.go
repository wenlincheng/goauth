package wechat

type CommonErr struct {
	ErrCode int    `json:"errcode"` // 错误码
	ErrMsg  string `json:"errmsg"`  // 错误信息
}

type AccessTokenResp struct {
	AccessToken  string `json:"access_token"`  // 接口调用凭证
	ExpiresIn    string `json:"expires_in"`    // 接口调用凭证超时时间 单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新 access_token
	Openid       string `json:"openid"`        // 授权用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
}

type UserInfoResp struct {
	Openid     string   `json:"openid"`     // 标识
	Unionid    string   `json:"unionid"`    // 统一标识
	Nickname   string   `json:"nickname"`   // 昵称
	HeadImgUrl string   `json:"headimgurl"` // 头像
	Sex        int      `json:"sex"`        // 性别
	Province   string   `json:"province"`   // 省份
	City       string   `json:"city"`       // 城市
	Country    string   `json:"country"`    // 国家 中国为 CN
	Privilege  []string `json:"privilege"`  // 特权信息
}
