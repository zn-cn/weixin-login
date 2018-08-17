package controller

import (
	"config"
	"errors"
	"fmt"
	"model"
	"net/http"
	"time"
	"util"
	"util/log"
	"util/res"

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
 *
 * @apiParamExample   {query} Request-Example:
 *    {
 *      "index_url": "https://weixin.bingyan-tech.hustonline.net/gradmovie/",
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
	conf := config.Conf
	url := fmt.Sprintf(conf.URL.CodeURL, conf.Wechat.AppID,
		conf.Wechat.RedirectURI, conf.Wechat.ResponseType, indexURL)
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

	indexURL := c.QueryParam("state")
	if indexURL == "" {
		writeWeixinLog("SetUserInfoByCode", "state is missing", errors.New("state is missing"))
		return res.RetError(http.StatusBadRequest, http.StatusBadRequest, "state is missing", c)
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
	userInfoMap := util.JSONStructToMap(userInfo)

	// set cookie
	for key, value := range userInfoMap {
		cookieValue := ""
		if v, ok := value.(string); ok {
			cookieValue = v
		}
		if key == "sex" {
			if v, ok := value.(float64); ok {
				cookieValue = string(int(v))
			}
		}
		if cookieValue != "" {
			cookie := new(http.Cookie)
			cookie.Name = key
			cookie.Value = cookieValue
			cookie.Expires = time.Now().Add(7 * 24 * time.Hour)
			cookie.Path = "/"
			cookie.HttpOnly = true // 必须
			c.SetCookie(cookie)
		}
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
