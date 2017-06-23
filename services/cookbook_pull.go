package services

import (
	//"sweetcook-backend/utils/request"
	"fmt"
	"sweetcook-backend/utils/logger"
	"reflect"
	"io/ioutil"
	"encoding/json"
	"sweetcook-backend/utils/mongodb"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"strconv"
	"sweetcook-backend/utils/request"
)

func init()  {
	//fmt.Println("run result: ", RunCookData())
}

const urlTemplate = "http://dy.api.duominuo.com/pdc?partnerId=ea86732b8e86388470ec8392ba949baa&token=9dcbc541377b8eb16f609df26c8d02a2&apiId=23&name=%E6%8E%92%E9%AA%A8&fields=&rn=rnfiled&pn=pnfield"

func RunCookData() (succ bool) {
	db := mongodb.ConnectMongo()
	pn,rn := 0, 500
	
	totalCount := 0
	insertCount := 0
	for pn < 1 {
	//for (pn+1) * rn < 16975 {
		reqUrl := generateReqUrl(urlTemplate, rn, pn)
		retJson, err := request.HttpGetJson(reqUrl)
		//retJson, err := FakeHttpGetJson(reqUrl)
		if err != nil {
			logger.Error("err: ", err)
			succ = false
			return
		}
		
		result := retJson["result"]
		resultMap := result.(map[string]interface{})
		data := resultMap["data"]
		dataMap := data.(map[string]interface{})
		dataList := dataMap["dataList"]
		dataItems := ToSlice(dataList)
		for _, dataItem := range dataItems {
			dataItemMap := dataItem.(map[string]interface{})
			cookBook := Cookbook{}
			cookBook.Adminicles = convertInterface(dataItemMap["adminicles"])
			cookBook.Category1 = convertInterface(dataItemMap["category1"])
			cookBook.Category2 = convertInterface(dataItemMap["category2"])
			cookBook.Content = convertInterface(dataItemMap["content"])
			cookBook.Intro = convertInterface(dataItemMap["intro"])
			cookBook.MainMaterials = convertInterface(dataItemMap["mainMaterials"])
			cookBook.Picture = convertInterface(dataItemMap["picture"])
			cookBook.Tips = convertInterface(dataItemMap["tips"])
			cookBook.Title = convertInterface(dataItemMap["title"])
			cookBook.Mid = convertInterface(dataItemMap["mid"])
			
			query := bson.M{"mid": cookBook.Mid}
			existedCookBook := Cookbook{}
			db.C(CollectionCookbook).Find(query).One(&existedCookBook)
			if existedCookBook.Mid == "" {
				//insert
				db.C(CollectionCookbook).Insert(cookBook)
				insertCount += 1
				logger.Debug("cookBook inserted : ", cookBook.Title)
			}else {
				logger.Debug("data existed. mid: ",existedCookBook.Mid, ", title: ", existedCookBook.Title, ", new title:", cookBook.Title)
			}
			totalCount += 1
		}
		pn += 1
	}
	logger.Debug("total insert count is: ", insertCount, ", and total count: ", totalCount)
	succ = true
	return
}

func generateReqUrl(url string, rn int, pn int) (reqUrl string) {
	reqUrl = strings.Replace(url, "rnfiled", strconv.Itoa(rn), -1)
	reqUrl = strings.Replace(reqUrl, "pnfield", strconv.Itoa(pn), -1)
	return
}

func convertInterface(param interface{}) (ret string) {
	ret = param.(string)
	return
}

func ToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		panic("toslice arr not slice")
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}

func FakeHttpGetJson(reqUrl string) (retMap map[string]interface{}, err error) {
	logger.Debug("requrl is: ", reqUrl)
	retMap,err = readFile("test/testcook.txt")
	return
}

func readFile(filename string) (retMap map[string]interface{}, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return
	}
	
	if err = json.Unmarshal(bytes, &retMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return
	}
	
	return
}