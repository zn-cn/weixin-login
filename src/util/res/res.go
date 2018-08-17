/*
Pakcage res response to the client with fixed format
*/
package res

import (
	"net/http"

	"github.com/labstack/echo"
)

// ErrorRes ErrorResponse
type ErrorRes struct {
	Status int    `json:"status"`
	ErrMsg string `json:"err_msg"`
}

// DataRes DataResponse
type DataRes struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// RetError response error, wrong response
func RetError(code, status int, errMsg string, c echo.Context) error {
	return c.JSON(code, ErrorRes{
		Status: status,
		ErrMsg: errMsg,
	})
}

// RetData response data, correct response
func RetData(data interface{}, c echo.Context) error {
	return c.JSON(http.StatusOK, DataRes{
		Status: 200,
		Data:   data,
	})
}
