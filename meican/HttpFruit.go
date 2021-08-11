package meican

import (
	"encoding/json"
	"fmt"
	"fruit-eudemon/rlog"
	"github.com/kirinlabs/HttpRequest"
)

type TokenResult struct {
	AccessToken       string `json:"access_token"`
	RefreshToken      string `json:"refresh_token"`
	ExpiresIn         int64  `json:"expires_in"`
	NeedResetPassword string `json:"need_reset_password"`
}

func DoTask() {
	rlog.Info("开始请求水果数据", nil)
	LoginCheck()
	rlog.Info("结束请求水果数据", nil)
}

func LoginCheck() {
	req := HttpRequest.NewRequest().SetTimeout(20)
	response, err := req.Post("https://meican.com/preference/preorder/api/v2.0/oauth/token", map[string]interface{}{
		"grant_type":             "password",
		"meican_credential_type": "password",
		"password":               "",
		"username":               "",
		"username_type":          "username",
		"remember":               "true",
		//需要携带clientId 否则后端会返回错误
		"client_id":     "Xqr8w0Uk4ciodqfPwjhav5rdxTaYepD",
		"client_secret": "vD11O6xI9bG3kqYRu9OyPAHkRGxLh4E",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Close()
	body, err := response.Body()
	if err != nil {
		return
	}
	var result = string(body)
	rlog.Info("返回值：", map[string]interface{}{
		"result": result,
	})

	var tokenResult TokenResult
	err2 := json.Unmarshal(body, &tokenResult)
	if err2 != nil {
		fmt.Println(err2)
	}

}
