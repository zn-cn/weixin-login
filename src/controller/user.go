package controller

import (
	"util/log"
	"util/res"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

var userLogger = log.GetLogger()

// GetUserInfo get user info
/**
 * @apiDefine GetUserInfo GetUserInfo
 * @apiDescription 获取当前用户信息
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 200,
 *       "data": {
 *           "nickname": String,
 *           "sex": "1"
 *           "province": String,
 *           "city": String,
 *           "country": "1"
 *           "headimgurl": String,
 *         }
 *     }
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
 * @api {get} /api/v1/user/info GetUserInfo
 * @apiVersion 1.0.0
 * @apiName GetUserInfo
 * @apiGroup User
 * @apiUse GetUserInfo
 */
func GetUserInfo(c echo.Context) error {
	cookies := c.Cookies()
	data := map[string]interface{}{}

	for _, value := range cookies {
		data[value.Name] = value.Value
	}
	// 去除敏感信息
	delete(data, "openid")
	delete(data, "unionid")
	return res.RetData(data, c)
}

func writeUserLog(funcName, errMsg string, err error) {
	userLogger.WithFields(logrus.Fields{
		"package":  "controller",
		"file":     "user.go",
		"function": funcName,
		"err":      err,
	}).Infoln(errMsg)
}
