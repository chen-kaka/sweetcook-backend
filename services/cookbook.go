package services

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
	"sweetcook-backend/utils/logger"
)

const (
	CollectionCookbook = "cookbook"
	CollectionJuheCookbook = "juhe_cookbook"
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
)

func QueryCookbookList(c *gin.Context, page int, num int) (cookbookList []JuheCookbook, err error) {
	db := c.MustGet("db").(*mgo.Database)
	query := bson.M{}
	logger.Debug("query: ", query, "page: ", page, "num: ", num)
	err = db.C(CollectionJuheCookbook).Find(query).Sort("id").Skip(page * num).Limit(num).All(&cookbookList)
	return
}