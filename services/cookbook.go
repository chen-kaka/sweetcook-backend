package services

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"sweetcook-backend/utils/logger"
	"time"
)

const (
	CollectionCookbook = "cookbook"
	CollectionJuheCookbook = "juhe_cookbook"
	CollectionUserCookbook = "user_cookbook"
)

type (
	Cookbook struct {
		Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		MainMaterials     string   `json:"mainMaterials" bson:"mainMaterials"`
		Adminicles     string   `json:"adminicles" bson:"adminicles"`
		Category1     string   `json:"category1" bson:"category1"`
		Category2     string   `json:"category2" bson:"category2"`
		Intro     string   `json:"intro" bson:"intro"`
		Picture     string   `json:"picture" bson:"picture"`
		Content     string   `json:"content" bson:"content"`
		Title     string   `json:"title" bson:"title"`
		Tips     string   `json:"tips" bson:"tips"`
		Mid     string   `json:"mid" bson:"mid"`
		CreatedAt int64         `json:"created_at" bson:"created_at"`
		UpdatedAt int64         `json:"updated_at" bson:"updated_at"`
	}
	
	JuheCookbook struct {
		Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Imtro     string   `json:"imtro" bson:"imtro"`
		Ingredients     string   `json:"ingredients" bson:"ingredients"`
		Tags     string   `json:"tags" bson:"tags"`
		Burden     string   `json:"burden" bson:"burden"`
		Albums     []interface{}   `json:"albums" bson:"albums"`
		Steps     []interface{}   `json:"steps" bson:"steps"`
		Title     string   `json:"title" bson:"title"`
		Mid      string     `json:"id" bson:"id"`
		CreatedAt int64         `json:"created_at" bson:"created_at"`
		UpdatedAt int64         `json:"updated_at" bson:"updated_at"`
	}
	
	UserCookbook struct {
		Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Username     string   `json:"username" bson:"username"`
		Ids []string  `json:"ids" bson:"ids" binding:"required"`
		CreatedAt int64         `json:"created_at" bson:"created_at"`
		UpdatedAt int64         `json:"updated_at" bson:"updated_at"`
	}
)

func QueryCookbookList(c *gin.Context, page int, num int) (cookbookList []JuheCookbook, err error) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{}
	logger.Debug("query: ", query, "page: ", page, "num: ", num)
	err = db.C(CollectionJuheCookbook).Find(query).Sort("id").Skip(page * num).Limit(num).All(&cookbookList)
	return
}

func CreateCookbook(c *gin.Context, userCookbook UserCookbook) (ret string, err error) {
	db := c.MustGet("db").(*mgo.Database)
	
	userCookbook.Id = bson.NewObjectId()
	userCookbook.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	
	logger.Debug("insert userCookbook: ", userCookbook)
	err = db.C(CollectionUserCookbook).Insert(userCookbook)
	if err != nil {
		logger.Error(err)
		return
	}
	ret = "create succ"
	return
}

func AddCookbookList(c *gin.Context, userCookbook UserCookbook) (retMsg string, err error) {
	db := c.MustGet("db").(*mgo.Database)
	
	existedUserCookbook := UserCookbook{}
	query := bson.M{"username": userCookbook.Username}
	db.C(CollectionUserCookbook).Find(query).One(&existedUserCookbook)
	
	if existedUserCookbook.Username == "" {
		return CreateCookbook(c, userCookbook)
	}
	
	logger.Debug("update")
	ids := existedUserCookbook.Ids
	for _,item := range userCookbook.Ids {
		existed := false
		for _, eItem := range ids {
			if eItem == item {
				existed = true
				break
			}
		}
		if !existed {
			ids = append(ids, item)
		}
	}
	existedUserCookbook.Ids = ids
	existedUserCookbook.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	
	err = db.C(CollectionUserCookbook).Update(query, existedUserCookbook)
	if err != nil {
		logger.Error(err)
		return
	}
	retMsg = "update succ"
	return
}

func QueryUserCookbooks(c *gin.Context, username string) (juheCookbooks []JuheCookbook, err error) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"username": username}
	logger.Debug("query cookbook: ", query)
	existedUserCookbook := UserCookbook{}
	err = db.C(CollectionUserCookbook).Find(query).One(&existedUserCookbook)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Debug("existedUserCookbook: ", existedUserCookbook)
	cookbookQuery := bson.M{"id": bson.M{"$in":existedUserCookbook.Ids}}
	logger.Debug("query cookbookQuery: ", cookbookQuery)
	err = db.C(CollectionJuheCookbook).Find(cookbookQuery).All(&juheCookbooks)
	return
}

func DeleteCookbookList(c *gin.Context, cookIds UserCookbook) (retMsg string, err error) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"username": cookIds.Username}
	logger.Debug("query cookbook: ", query)
	existedUserCookbook := UserCookbook{}
	err = db.C(CollectionUserCookbook).Find(query).One(&existedUserCookbook)
	if err != nil {
		logger.Error(err)
		return
	}
	
	newIds := []string{}
	for _,eItem := range existedUserCookbook.Ids {
		existed := false
		for _, item := range cookIds.Ids {
			if eItem == item {
				existed = true
				break
			}
		}
		if !existed {
			newIds = append(newIds, eItem)
		}
	}
	existedUserCookbook.Ids = newIds
	existedUserCookbook.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	
	err = db.C(CollectionUserCookbook).Update(query, existedUserCookbook)
	if err != nil {
		logger.Error(err)
		return
	}
	retMsg = "delete succ"
	return
}