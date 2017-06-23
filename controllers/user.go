package controllers

import (
	"github.com/gin-gonic/gin"
	"sweetcook-backend/utils/error"
	"sweetcook-backend/utils/logger"
	"sweetcook-backend/services"
	"github.com/gin-gonic/contrib/sessions"
	"time"
)

type (
	User struct {
	}
)

/**
http://localhost:7000/app-api/user/info

 */
//获取数据
func (u User)Info(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	userInfo,err := services.QueryOne(c, sessionUsername)
	if err != nil {
		logger.Error(err)
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_DATABASE_ERROR,error.MSG_NAME: err})
		return
	}
	
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.DATA_NAME: userInfo})
}

/*
设置数据

http://localhost:7000/app-api/user/register

{"username":"kaka","telephone":"1234567890","nickname":"kakachan","password":"123456","sex":1}
 */

func (u User)Register(c *gin.Context)  {
	var userInfo services.UserInfo
	if err := c.BindJSON(&userInfo); err != nil {
		logger.Error(err)
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	savedUserInfo,err := services.CreateOrUpdate(c, userInfo, false)
	if err != nil {
		logger.Error(err)
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_DATABASE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: savedUserInfo})
}

/*
	登录

http://localhost:7000/app-api/user/login

{"username":"kaka","password":"123456"}
 */

func (u User)Login(c *gin.Context)  {
	var loginInfo services.LoginInfo
	if err := c.BindJSON(&loginInfo); err != nil {
		logger.Error(err)
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	
	userInfo,err := services.QueryOne(c, loginInfo.Username)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "user not exist."})
		return
	}
	logger.Debug("user info: ", userInfo)
	if userInfo.Password == loginInfo.Password {
		session := sessions.Default(c)
		session.Set("username", userInfo.Username)
		session.Set("login_at", time.Now().String())
		session.Set("nickName", userInfo.Nickname)
		session.Save()
		c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: "login succ."})
		return
	}
	
	c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_LOGIN_FAILED,error.MSG_NAME: "login failed."})
}

/*
	退出登录

http://localhost:7000/app-api/user/logout

 */

func (u User)Logout(c *gin.Context)  {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: "Signed out"})
}

/**
	绑定情侣账号功能
	
	http://localhost:7000/app-api/user/bind?bind_username=xy
 */
func (u User)Bind(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	bindUsername := c.Query("bind_username")
	
	bindUserInfo,err := services.QueryOne(c, bindUsername)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "目标用户名不存在。"})
		return
	}
	
	userInfo,err := services.QueryOne(c, sessionUsername)
	
	if bindUserInfo.Companion != "" {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "目标用户已绑定情侣。"})
		return
	}
	
	retMsg,err := services.BindUser(c, userInfo, bindUserInfo)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
	}
	
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: retMsg})
}

func GetUserInfo(c *gin.Context) (username string) {
	//拿出session
	session := sessions.Default(c)
	sessionUsername := session.Get("username")
	logger.Debug("sessionUsername: ", sessionUsername)
	if sessionUsername == nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "user not login."})
		return
	}
	
	username,_ = sessionUsername.(string)
	return
}