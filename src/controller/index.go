package controller

import (
	"util/res"

	"github.com/labstack/echo"
)

// CheckHealthy CheckHealthy
/**
 * @apiDefine CheckHealthy CheckHealthy
 * @apiDescription CheckHealthy
 *
 * @apiSuccess {Number} status=200 状态码
 * @apiSuccess {Object} data 正确返回数据
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "status": 200,
 *       "data": "hello world"
 *     }
 *
 * @apiError {Number} status 状态码
 * @apiError {String} err_msg 错误信息
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 500 GateWay
 *     {
 *       "status": 500,
 *       "err_msg": "server error"
 *     }
 */
/**
 * @api {get} /api/v1/health CheckHealthy
 * @apiVersion 1.0.0
 * @apiName CheckHealthy
 * @apiGroup Index
 * @apiUse CheckHealthy
 */
func CheckHealthy(c echo.Context) error {
	return res.RetData("hello world", c)
}
