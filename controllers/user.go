package controllers

import (
	"github.com/gin-gonic/gin"
	"sweetcook-backend/utils/error"
	"sweetcook-backend/utils/logger"
	"sweetcook-backend/services"
)

type (
	User struct {
	}
)

/**
http://localhost:7000/app-api/user/info?username=kaka

 */
//获取数据
func (u User)Info(c *gin.Context)  {
	var username = c.Query("username")
	if username == "" {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "username is null."})
		return
	}
	
	userInfo,err := services.QueryOne(c, username)
	if err != nil {
		logger.Error(err)
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_DATABASE_ERROR,error.MSG_NAME: err.Error()})
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
	savedUserInfo,err := services.CreateOrUpdate(c, userInfo)
	if err != nil {
		logger.Error(err)
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_DATABASE_ERROR,error.MSG_NAME: err.Error()})
	}
	
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: savedUserInfo})
}