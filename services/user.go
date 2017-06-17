package services

import (
	"time"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sweetcook-backend/utils/logger"
)

const (
	CollectionUser = "user"
)

type (
	UserInfo struct {
		Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Username     string   `json:"username" bson:"username" binding:"required"`
		Nickname      string        `json:"nickname" bson:"nickname" binding:"required"`
		Password      string        `json:"password" bson:"password" binding:"required"`
		Telephone     string    `json:"telephone" bson:"telephone"`
		Sex           int    `json:"sex" bson:"sex" binding:"required"` // 0: male 1: female
		Companion    bson.ObjectId         `json:"companion_id" bson:"companion_id"`
		CreatedAt int64         `json:"created_at" bson:"created_at"`
		UpdatedAt int64         `json:"updated_at" bson:"updated_at"`
	}
)

func Create(c *gin.Context, userInfo UserInfo) (ret string, err error) {
	db := c.MustGet("db").(*mgo.Database)
	
	userInfo.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	
	err = db.C(CollectionUser).Insert(userInfo)
	if err != nil {
		logger.Error(err)
		return
	}
	ret = "create succ"
	return
}

// List all userConfigs
func QueryOne(c *gin.Context, username string) (userInfo UserInfo, err error) {
	db := c.MustGet("db").(*mgo.Database)

	query := bson.M{"username": username}
	err = db.C(CollectionUser).Find(query).One(&userInfo)
	return
}

// Update an userConfig
func CreateOrUpdate(c *gin.Context, userInfo UserInfo) (ret interface{}, err error) {
	db := c.MustGet("db").(*mgo.Database)
	
	if userInfo.Id == "" {
		userInfo.Id = bson.NewObjectId()
	}
	logger.Debug("id: ", userInfo.Id)
	
	existedUserInfo := UserInfo{}
	query := bson.M{"username": userInfo.Username}
	db.C(CollectionUser).Find(query).One(&existedUserInfo)
	
	if existedUserInfo.Username == "" {
		return Create(c, userInfo)
	}
	
	logger.Debug("update")
	userInfo.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	
	err = db.C(CollectionUser).Update(query, userInfo)
	if err != nil {
		logger.Error(err)
		return
	}
	ret = "update succ"
	return
}