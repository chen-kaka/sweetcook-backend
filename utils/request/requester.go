package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sweetcook-backend/utils/logger"
	"encoding/json"
	"bytes"
)


func HttpGetJson(reqUrl string) ( retJson map[string]interface{}, err error ) {
	body,err := HttpGet(reqUrl)
	err = json.Unmarshal(body, &retJson)
	logger.Debug("resp is: ", retJson)
	return
}

func HttpGet(reqUrl string) ( retBytes []byte, err error ) {
	fmt.Println("req url is: ", reqUrl)
	response, err := http.Get(reqUrl)
	if err != nil {
		logger.Error(err)
		return
	}
	defer response.Body.Close()
	retBytes, err = ioutil.ReadAll(response.Body)
	return
}

func HttpPostAllJson(reqUrl string, postContent map[string]interface{}) ( retJson map[string]interface{}, err error ) {
	body, err := HttpPostJson(reqUrl, postContent)
	err = json.Unmarshal(body, &retJson)
	logger.Debug("resp is: ", retJson)
	return
}

func HttpPostJson(reqUrl string, postContent map[string]interface{}) ( retBytes []byte, err error ) {
	fmt.Println("req url is: ", reqUrl)
	
	contentByte, err := json.Marshal(postContent)
	reqBody := bytes.NewBuffer(contentByte)
	response, err := http.Post(reqUrl, "application/json", reqBody)
	if err != nil {
		logger.Error(err)
		return
	}
	
	defer response.Body.Close()
	retBytes, err = ioutil.ReadAll(response.Body)
	return
}

func HttpPostForm(reqUrl string, postContent map[string]string) ( retBytes []byte, err error ) {
	fmt.Println("req url is: ", reqUrl)
	
	queryString := mapToQueryString(postContent)
	reqBody := bytes.NewBuffer([]byte(queryString))
	response, err := http.Post(reqUrl, "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		logger.Error(err)
		return
	}
	
	defer response.Body.Close()
	retBytes, err = ioutil.ReadAll(response.Body)
	return
}

func mapToQueryString(m map[string]string) string {
	paramsStr := ""
	for k,v := range m {
		paramsStr += fmt.Sprintf("%s=%s&", k ,v)
	}
	return paramsStr
}