package services

import "gopkg.in/mgo.v2/bson"

type Cookbook struct {
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
	CreatedAt int64         `json:"created_at" bson:"created_at"`
	UpdatedAt int64         `json:"updated_at" bson:"updated_at"`
}