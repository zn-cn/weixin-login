package model

import (
	"config"
	"util/log"

	"github.com/imroc/req"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

// UserInfo 用户信息
type UserInfo struct {
	Openid     string   `json:"openid" bson:"openid"`
	Nickname   string   `json:"nickname" bson:"nickname"`
	Sex        int      `json:"sex" bson:"sex"`
	Province   string   `json:"province" bson:"province"`
	City       string   `json:"city" bson:"city"`
	Country    string   `json:"country" bson:"country"`
	HeadImgURL string   `json:"headimgurl" bson:"headimgurl"`
	Privilege  []string `json:"privilege" bson:"privilege"`
	Unionid    string   `json:"unionid" bson:"unionid"`
	Errcode    int      `json:"errcode" bson:"errcode"`
	Errmsg     string   `json:"errmsg" bson:"errmsg"`
}

// WeixinTokenRes 微信 access_token 获取返回数据结构体
type WeixinTokenRes struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
}

var weixinLogger = log.GetLogger()

// GetWeixinAccessToken get the access token of WeChat Official Account
func GetWeixinAccessToken(code string) (WeixinTokenRes, error) {
	conf := config.Conf
	param := req.Param{
		"appid":      conf.Wechat.AppID,
		"secret":     conf.Wechat.AppSecret,
		"code":       code,
		"grant_type": "authorization_code",
	}

	weixinTokenRes := WeixinTokenRes{}
	err := BindGetJSONData(conf.URL.TokenURL, param, &weixinTokenRes)
	if err != nil {
		writeWeixinLog("GetWeixinJSAPI", "get the weixin access_token error", err)
		return WeixinTokenRes{}, err
	}
	return weixinTokenRes, nil
}

// GetUserInfo get the userInfo by weixin access_token and openid
func GetUserInfo(accessToken, openid string) (UserInfo, error) {
	param := req.Param{
		"access_token": accessToken,
		"openid":       openid,
		"lang":         "zh_CN",
	}
	userInfo := UserInfo{}
	err := BindGetJSONData(config.Conf.URL.UserInfoURL, param, &userInfo)
	if err != nil {
		writeWeixinLog("GetUserInfo", "get the weixin access_token error", err)
		return UserInfo{}, err
	}
	return userInfo, nil
}

func SetRedisURL(key, value string) error {
	ctrl := NewRedisDBCntlr()
	defer ctrl.Close()
	_, err := ctrl.SETEX(key, 10*60, value)
	return err
}

func GetRedisURL(key string) (map[string]string, error) {
	ctrl := NewRedisDBCntlr()
	defer ctrl.Close()
	value, _ := ctrl.GET(key)
	res := map[string]string{}
	err := jsoniter.UnmarshalFromString(value, &res)
	return res, err
}

func writeWeixinLog(funcName, errMsg string, err error) {
	weixinLogger.WithFields(logrus.Fields{
		"package":  "model",
		"file":     "weixin.go",
		"function": funcName,
		"err":      err,
	}).Infoln(errMsg)
}
