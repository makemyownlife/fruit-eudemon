package meican

import (
	"fmt"
	"fruit-eudemon/rlog"
	"github.com/kirinlabs/HttpRequest"
)

func DoTask() {
	rlog.Info("开始请求水果数据", nil)
	LoginCheck()
	rlog.Info("结束请求水果数据", nil)
}

func LoginCheck() {
	req := HttpRequest.NewRequest().SetTimeout(20)
	response, err := req.Get("http://172.31.132.190:18080/consumer/groupList.query")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Close()
	body, err := response.Body()
	if err != nil {
		return
	}
	rlog.Info("返回值:", map[string]interface{}{
		"result": string(body),
	})
}
