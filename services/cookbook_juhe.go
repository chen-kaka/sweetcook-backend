package services

import (
	//"sweetcook-backend/utils/request"
	"fmt"
	"sweetcook-backend/utils/logger"
	"sweetcook-backend/utils/mongodb"
	"gopkg.in/mgo.v2/bson"
	//"sweetcook-backend/utils/request"
	"strings"
	"strconv"
	"sweetcook-backend/utils/request"
)

func init()  {
	fmt.Println("run result: ", RunJuheCookData())
}

const juheUrlTemplate = "http://apis.juhe.cn/cook/index?key=da3e5f2014978b8617b4d0c1f8c169ce&cid=cidfield&rn=rnfiled&pn=pnfield"

func RunJuheCookData() (succ bool) {
	db := mongodb.ConnectMongo()
	pn,rn := 0, 30
	
	totalCount := 0
	insertCount := 0
	
	for cid := 350;cid <= 360;cid += 1 {
		logger.Debug("cid is: ", cid)
		pn = 0
		for pn < 1 {
			reqUrl := generateJuheReqUrl(juheUrlTemplate, rn, pn, cid)
			retJson, err := request.HttpGetJson(reqUrl)
			//retJson, err := FakeJuheHttpGetJson(reqUrl)
			if err != nil {
				logger.Error("err: ", err)
				succ = false
				return
			}
			
			result := retJson["result"]
			if result == nil {
				logger.Error("process end!.")
				continue
			}
			resultMap := result.(map[string]interface{})
			if resultMap == nil {
				logger.Error("process end!.")
				break
			}
			data := resultMap["data"]
			dataItems := ToSlice(data)
			for _, dataItem := range dataItems {
				dataItemMap := dataItem.(map[string]interface{})
				cookBook := JuheCookbook{}
				
				cookBook.Tags = convertInterface(dataItemMap["tags"])
				cookBook.Ingredients = convertInterface(dataItemMap["ingredients"])
				cookBook.Burden = convertInterface(dataItemMap["burden"])
				cookBook.Title = convertInterface(dataItemMap["title"])
				cookBook.Mid = convertInterface(dataItemMap["id"])
				
				stepsInterface := dataItemMap["steps"]
				if stepsInterface != nil {
					stepList := stepsInterface.([]interface{})
					cookBook.Steps = stepList
				}
				
				albumsInterface := dataItemMap["albums"]
				if albumsInterface != nil {
					albumList := albumsInterface.([]interface{})
					cookBook.Albums = albumList
				}
				
				query := bson.M{"id": cookBook.Mid}
				existedCookBook := JuheCookbook{}
				db.C(CollectionJuheCookbook).Find(query).One(&existedCookBook)
				if existedCookBook.Mid == "" {
					//insert
					db.C(CollectionJuheCookbook).Insert(cookBook)
					insertCount += 1
					logger.Debug("cookBook inserted : ", cookBook.Title)
				} else {
					logger.Debug("data existed. mid: ", existedCookBook.Mid, ", title: ", existedCookBook.Title, ", new title:", cookBook.Title)
				}
				totalCount += 1
			}
			pn += 1
		}
	}
	logger.Debug("total insert count is: ", insertCount, ", and total count: ", totalCount)
	succ = true
	return
}


func FakeJuheHttpGetJson(reqUrl string) (retMap map[string]interface{}, err error) {
	logger.Debug("requrl is: ", reqUrl)
	retMap,err = readFile("test/juhe_cookbook.txt")
	return
}

func generateJuheReqUrl(url string, rn int, pn int, cid int) (reqUrl string) {
	reqUrl = strings.Replace(url, "rnfiled", strconv.Itoa(rn), -1)
	reqUrl = strings.Replace(reqUrl, "pnfield", strconv.Itoa(pn), -1)
	reqUrl = strings.Replace(reqUrl, "cidfield", strconv.Itoa(cid), -1)
	return
}