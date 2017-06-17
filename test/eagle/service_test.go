package eagle

import (
	"testing"
	"sweetcook-backend/utils/eagle"
	"fmt"
	"sweetcook-backend/utils"
)

func Test_Get(t *testing.T)  {
	params := make(map[string]string)
	params["hello"] = "get"
	paramsStr := ""
	for k,v := range params {
		paramsStr += fmt.Sprintf("%s=%s&",k , v)
	}
	resp,err := eagle.Get("/v1/serve/provide", "level2", "ginshiba", "1.0.0", paramsStr)
	if err != nil {
		t.Error(err)
		return
	}
	
	respJson, err := utils.ParseBytesToJson(resp)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(respJson)
	t.Log("测试通过！")
}

func Test_PostJson(t *testing.T)  {
	params := make(map[string]interface{})
	params["hello"] = "postjson"
	resp,err := eagle.PostJson("/v1/serve/provide_post_json", "level2", "ginshiba", "1.0.0", params)
	if err != nil {
		t.Error(err)
		return
	}

	respJson, err := utils.ParseBytesToJson(resp)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(respJson)
	t.Log("测试通过！")
}

func Test_PostForm(t *testing.T)  {
	params := make(map[string]string)
	params["hello"] = "postform"
	resp,err := eagle.PostForm("/v1/serve/provide_post_form", "level2", "ginshiba", "1.0.0", params)
	if err != nil {
		t.Error(err)
		return
	}

	respJson, err := utils.ParseBytesToJson(resp)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(respJson)
	t.Log("测试通过！")
}