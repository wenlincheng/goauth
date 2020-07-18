package qq

import "github.com/wenlincheng/letigo/util"

/**

QQ

*/
type QQ struct {
	AppID     string
	AppSecret string
}

func GetQQClient() *QQ {

	util.GenRandNo(123)

	return new(QQ)
}
