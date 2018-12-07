package model

import (
	"github.com/imroc/req"
)

// BindGetJSONData bind the json data of method GET
// body must be a point
func BindGetJSONData(url string, param req.Param, body interface{}) error {
	r, err := req.Get(url, param)
	if err != nil {
		return err
	}
	err = r.ToJSON(body)
	if err != nil {
		return err
	}
	return nil
}

func ReqPOST(url string, body interface{}) (*req.Resp, error) {
	return req.Post(url, req.BodyJSON(body))
}
