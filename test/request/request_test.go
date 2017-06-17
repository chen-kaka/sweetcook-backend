package request

import (
	"testing"
	"sweetcook-backend/utils/request"
	"fmt"
	"encoding/json"
)

func Test_HttpPostAllJson(t *testing.T)  {
	var postContent = make(map[string]interface{})
	postContent["hello"] = "testalljson"
	retJson, err := request.HttpPostAllJson("http://10.2.124.15:12505/v1/serve/provide_post_json", postContent)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(retJson)
	t.Log("测试通过！")
}

func Test_HttpPostJson(t *testing.T)  {
	var postContent = make(map[string]interface{})
	postContent["hello"] = "testjson"
	retBytes, err := request.HttpPostJson("http://10.2.124.15:12505/v1/serve/provide_post_json", postContent)
	if err != nil {
		t.Error(err)
		return
	}
	var retJson map[string]interface{}
	err = json.Unmarshal(retBytes, &retJson)
	fmt.Println(retJson)
	t.Log("测试通过！")
}

func Test_HttpPostForm(t *testing.T)  {
	var postContent = make(map[string]string)
	postContent["hello"] = "testform"
	retBytes, err := request.HttpPostForm("http://10.2.124.15:12505/v1/serve/provide_post_form", postContent)
	if err != nil {
		t.Error(err)
		return
	}
	var retJson map[string]interface{}
	err = json.Unmarshal(retBytes, &retJson)
	fmt.Println(retJson)
	t.Log("测试通过！")
}