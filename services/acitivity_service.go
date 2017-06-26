package services

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"sweetcook-backend/utils/logger"
)

const (
	CollectionActivity = "activity"
)

type (
	Activity struct {
		Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Username     string   `json:"username" bson:"username"`
		CookIds     []string  `json:"cook_ids" bson:"cook_ids" binding:"required"`
		CookUser    string   `json:"cook_user" bson:"cook_user"`
		Content    string   `json:"content" bson:"content"`
		Comments    []Comment   `json:"comments" bson:"comments"`
		CreatedAt int64         `json:"created_at" bson:"created_at"`
		UpdatedAt int64         `json:"updated_at" bson:"updated_at"`
	}
	
	Comment struct {
		Id        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
		Username     string   `json:"username" bson:"username" binding:"required"`
		Comment    string   `json:"comment" bson:"comment" binding:"required"`
		CreatedAt int64         `json:"created_at" bson:"created_at"`
		UpdatedAt int64         `json:"updated_at" bson:"updated_at"`
	}
)

func CreateOrUpdateActivity(c *gin.Context, activity Activity) (retMsg string, err error) {
	db := c.MustGet("db").(*mgo.Database)
	if activity.Id.Hex() != "" {
		existedActivity := Activity{}
		query := bson.M{"_id": activity.Id}
		db.C(CollectionActivity).Find(query).One(&existedActivity)
		//更新
		activity.Comments = existedActivity.Comments
		activity.UpdatedAt = time.Now().UnixNano() / int64(time.Millisecond)
		activity.CreatedAt = existedActivity.CreatedAt
		
		err = db.C(CollectionActivity).Update(query, activity)
		if err != nil {
			logger.Error(err)
			return
		}
		retMsg = "update succ"
		return
	}
	return CreateActivity(c, activity)
}


func CreateActivity(c *gin.Context, activity Activity) (ret string, err error) {
	db := c.MustGet("db").(*mgo.Database)
	
	activity.Id = bson.NewObjectId()
	activity.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	
	logger.Debug("insert activity: ", activity)
	err = db.C(CollectionActivity).Insert(activity)
	if err != nil {
		logger.Error(err)
		return
	}
	ret = "create succ"
	return
}

func QueryActivityList(c *gin.Context, username string, page int, num int) (activityList []Activity, err error) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"username": username}
	logger.Debug("query: ", query, "page: ", page, "num: ", num)
	err = db.C(CollectionActivity).Find(query).Sort("id").Skip(page * num).Limit(num).All(&activityList)
	return
}

func DeleteActivity(c *gin.Context, username string, id string) (retMsg string, err error) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{"_id": bson.ObjectIdHex(id),"username":username}
	
	logger.Debug("query info is: ", query)
	err = db.C(CollectionActivity).Remove(query)
	retMsg = "delete succ."
	return
}