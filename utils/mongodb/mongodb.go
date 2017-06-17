package mongodb

import (
	"gopkg.in/mgo.v2"
	"sweetcook-backend/config"
	"sweetcook-backend/utils"
	"sweetcook-backend/utils/logger"
	"github.com/gin-gonic/gin"
)

var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)


func init()  {
	configJson := config.GetConfigJson()
	mongodbInfo := configJson["mongodb"]
	
	uri, err := utils.GetStringFromInterfaceMap(mongodbInfo, "uri")
	
	if err != nil {
		logger.Error(err)
		logger.Error("mongodb config uri not found. exit.")
		return
	}
	
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		logger.Error("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	logger.Info("Connected to", uri)
	Session = s
	Mongo = mongo
}

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	s := Session.Clone()
	
	defer s.Close()
	
	c.Set("db", s.DB(Mongo.Database))
	c.Next()
}

func ConnectMongo()  *mgo.Database {
	s := Session.Clone()
	return s.DB(Mongo.Database)
}