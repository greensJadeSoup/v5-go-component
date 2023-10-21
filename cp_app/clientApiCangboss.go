package cp_app

import (
	"v5-go-component/cp_api"
	"v5-go-component/cp_constant"
	"v5-go-component/cp_error"
	"v5-go-component/cp_obj"
)

// 调用sso session信息获取接口
func CheckSession(solder *BaseController) (*cp_api.CheckSessionInfo, error) {
	respBody, err := Instance.CallClient.NewCall(solder, cp_api.SVRAPI_SSO_SESSION_CHECK, "")
	if err != nil {
		return nil, err
	}

	respObj := &struct {
		Code    int
		Message string
		Data    cp_api.CheckSessionInfo
	}{}

	err = cp_obj.Cjson.Unmarshal(respBody, respObj)
	if err != nil {
		return nil, cp_error.NewNormalError(err)
	} else if respObj.Code != cp_constant.RESPONSE_CODE_OK {
		return nil, cp_error.NewNormalError("获取session信息请求失败:" + respObj.Message)
	}

	return &respObj.Data, nil
}
