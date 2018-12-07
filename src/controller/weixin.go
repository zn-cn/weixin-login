package controller

import (
	"config"
	"errors"
	"fmt"
	"model"
	"net/http"
	"strings"
	"time"
	"util/log"
	"util/res"

	"github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"

	"net/url"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var weixinLogger = log.GetLogger()

// GetRedirectURI get the weixin redirect_uri
/**
 * @apiDefine GetRedirectURI GetRedirectURI
 * @apiDescription get the weixin redirect_uri
 *
 * @apiParam  {String} index_url 首页URL，后台跳转回首页的URL
 * @apiParam  {String} cb_api 回调API接口：即为之后后台拿到用户信息之后后台会调用此接口将API发送过去，限制<another host>和<host>的子域名, method:POST
 *
 * @apiParamExample   {query} Request-Example:
 *    {
 *      "index_url": "https://<weixin host>/gradmovie/",
 *      "cb_api": "https://test.<host>/api/v1/userInfo"
 *    }
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 200,
 *       "data": {
 *           "redirect_uri": "https://open.weixin.qq.com/connect/oauth2/authorize?appid=APPID&redirect_uri=REDIRECT_URI&response_type=code&scope=SCOPE&state=STATE#wechat_redirect"
 *         }
 *     }
 *
 * @apiError {Number} status 状态码
 * @apiError {String} err_msg 错误信息
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 400 BadRequest
 *     {
 *       "status": 400,
 *       "err_msg": "lack the param:index_url"
 *     }
 */
/**
 * @api {get} /api/v1/weixin/redirect_uri GetRedirectURI
 * @apiVersion 1.0.0
 * @apiName GetRedirectURI
 * @apiGroup Weixin
 * @apiUse GetRedirectURI
 */
func GetRedirectURI(c echo.Context) error {
	indexURL := c.QueryParam("index_url")
	if indexURL == "" {
		writeWeixinLog("GetRedirectURI", "lack the param:index_url", errors.New("lack the param:index_url"))
		return res.RetError(http.StatusBadRequest, http.StatusBadRequest, "lack the param:index_url", c)
	}
	state := map[string]string{
		"index_url": indexURL,
	}
	cbAPI := c.QueryParam("cb_api")
	if cbAPI != "" {
		u, err := url.Parse(cbAPI)
		if err != nil {
			writeWeixinLog("GetRedirectURI", "lack the param:cb_api", errors.New("lack the param:cb_api"))
			return res.RetError(http.StatusBadRequest, http.StatusBadRequest, "lack the param:cb_api", c)
		}
		hostname := u.Hostname()
		hostnameSplit := strings.Split(hostname, ".")

		valid := false
		for _, h := range config.Conf.URL.WhiteList {
			hSplit := strings.Split(h, ".")
			tempValid := true
			for i := 0; i < len(hSplit); i++ {
				if hSplit[len(hSplit)-i-1] != hostnameSplit[len(hostnameSplit)-i-1] {
					tempValid = false
					break
				}
			}
			if tempValid {
				valid = true
				break
			}
		}
		if !valid {
			writeWeixinLog("GetRedirectURI", "cb_api 不在白名单内", errors.New("cb_api 不在白名单内"))
			return res.RetError(http.StatusBadRequest, http.StatusBadRequest, "cb_api 不在白名单内", c)
		}
		state["cb_api"] = cbAPI
	}
	stateStr, _ := jsoniter.MarshalToString(state)
	key := uuid.NewV4().String()

	err := model.SetRedisURL(key, stateStr)
	if err != nil {
		writeWeixinLog("GetRedirectURI", "redis操作出现问题", errors.New("redis操作出现问题"))
		return res.RetError(http.StatusBadGateway, http.StatusBadGateway, "redis操作出现问题", c)
	}

	conf := config.Conf
	url := fmt.Sprintf(conf.URL.CodeURL, conf.Wechat.AppID,
		conf.Wechat.RedirectURI, conf.Wechat.ResponseType, key)
	data := map[string]interface{}{
		"redirect_uri": url,
	}
	return res.RetData(data, c)
}

// SetUserInfoByCode get wechat access
/**
 * @apiDefine SetUserInfoByCode SetUserInfoByCode
 * @apiDescription 通过微信 code 获取用户信息并设置cookie之后进行302跳转
 *
 * @apiParam  {String} code 微信code
 * @apiParam  {String} state index url
 *
 * @apiParamExample   {query} Request-Example:
 *    {
 *      "code": "code",
 *      "state": "index url"
 *    }
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 302 StatusFound
 *
 * @apiError {Number} status 状态码
 * @apiError {String} err_msg 错误信息
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 401 Unauthorized
 *     {
 *       "status": 401,
 *       "err_msg": "Unauthorized"
 *     }
 */
/**
 * @api {get} /api/v1/weixin/code SetUserInfoByCode
 * @apiVersion 1.0.0
 * @apiName SetUserInfoByCode
 * @apiGroup Weixin
 * @apiUse SetUserInfoByCode
 */
func SetUserInfoByCode(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		writeWeixinLog("SetUserInfoByCode", "code is missing", errors.New("code is missing"))
		return res.RetError(http.StatusBadRequest, http.StatusBadRequest, "code is missing", c)
	}

	state := c.QueryParam("state")
	data, err := model.GetRedisURL(state)
	if err != nil {
		writeWeixinLog("GetWeixinAccess", "get userInfo faild", err)
		return c.Redirect(http.StatusFound, "https://www.<another host>")
	}
	indexURL := data["index_url"]
	cbAPI := data["cb_api"]
	if indexURL == "" {
		writeWeixinLog("SetUserInfoByCode", "state is missing", errors.New("state is missing"))
		return c.Redirect(http.StatusFound, "https://www.<another host>")
	}
	weixinTokenRes, err := model.GetWeixinAccessToken(code)
	if err != nil {
		writeWeixinLog("GetWeixinAccess", "get weixin accessTokenRes faild", err)
		return c.Redirect(http.StatusFound, indexURL)
	}

	userInfo, err := model.GetUserInfo(weixinTokenRes.AccessToken, weixinTokenRes.Openid)
	if err != nil {
		writeWeixinLog("GetWeixinAccess", "get userInfo faild", err)
		return c.Redirect(http.StatusFound, indexURL)
	}
	// set openid cookie
	cookie := new(http.Cookie)
	cookie.Name = "openid"
	cookie.Value = userInfo.Openid
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
	cookie.Path = "/"
	cookie.HttpOnly = true // 必须
	c.SetCookie(cookie)

	if userInfo.Unionid != "" {
		cookie := new(http.Cookie)
		cookie.Name = "unionid"
		cookie.Value = userInfo.Unionid
		cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
		cookie.Path = "/"
		cookie.HttpOnly = true // 必须
		c.SetCookie(cookie)
	}

	if cbAPI != "" {
		go model.ReqPOST(cbAPI, userInfo)
	}

	return c.Redirect(http.StatusFound, indexURL)
}

func writeWeixinLog(funcName, errMsg string, err error) {
	weixinLogger.WithFields(logrus.Fields{
		"package":  "controller",
		"file":     "weixin.go",
		"function": funcName,
		"err":      err,
	}).Infoln(errMsg)
}
