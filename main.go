package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"sweetcook-backend/config"
	"sweetcook-backend/utils/logger"
	"sweetcook-backend/controllers"
	"sweetcook-backend/utils"
	"fmt"
	"net/http"
	"sweetcook-backend/utils/mongodb"
)

func main() {
	defer func() {
		err := recover()
		if err != nil {
			logger.Fatal(err)
			logger.Fatal("The ice breaks!")
		}
	}()

	r := gin.Default()

	configJson := config.GetConfigJson()
	sessionInfo := configJson["session"]
	useRedis,_ := utils.GetFloat64FromInterfaceMap(sessionInfo, "useRedis")
	if useRedis == 1 {
		redisHostPort,_ := utils.GetStringFromInterfaceMap(sessionInfo, "redisHostPort")
		redisSessionName,_ := utils.GetStringFromInterfaceMap(sessionInfo, "redisSessionName")
		redisRequirePass,_ := utils.GetStringFromInterfaceMap(sessionInfo, "redisRequirePass")
		logger.Info("using redis session， host: ", redisHostPort, ", session name: ", redisSessionName, ", redisRequirePass: ", redisRequirePass)
		//session
		store, _ := sessions.NewRedisStore(10, "tcp", redisHostPort, redisRequirePass, []byte("secret"))
		//store := sessions.NewCookieStore([]byte("secret"))
		r.Use(sessions.Sessions("mysession", store))

	} else {
		logger.Info("redis session not activate.")
	}
	//cors
	r.Use(CORSMiddleware())
	r.Use(mongodb.Connect)
	
	v1 := r.Group("/app-api")
	{
		userInfo := new(controllers.User)
		v1.GET("/user/info", userInfo.Info)
		v1.POST("/user/register", userInfo.Register)
		v1.POST("/user/login", userInfo.Login)
		v1.GET("/user/logout", userInfo.Logout)
		v1.GET("/user/bind", userInfo.Bind)
		
		cookbook := new(controllers.Cookbook)
		v1.GET("/cookbook/list", cookbook.List)
		v1.POST("/cookbook/add", cookbook.AddCookbookList)
		v1.GET("/cookbook/user_cookbooks", cookbook.QueryUserCookbooks)
		v1.GET("/cookbook/companion_cookbooks", cookbook.QueryCompanionCookbooks)
		v1.POST("/cookbook/delete", cookbook.DeleteCookbookList)
		
		acitivity := new(controllers.Activity)
		v1.POST("/activity/create_update", acitivity.CreateOrUpdateActivity)
		v1.GET("/activity/list", acitivity.ListActivity)
		v1.GET("/activity/delete", acitivity.DeleteActivity)
		v1.POST("/activity/comment_rate", acitivity.CommentAndRateActivity)
		v1.GET("/activity/set_finished", acitivity.SetActivityFinished)
		v1.GET("/activity/single_dog", acitivity.SingleDogActivityInfo)
	}

	r.Static("/public", "./public")
	r.LoadHTMLGlob("public/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"hint": "it works!"})
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	
	r.Run(":7000") // listen and serve on 0.0.0.0:7000
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		
		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}