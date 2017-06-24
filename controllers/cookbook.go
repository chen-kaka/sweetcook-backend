package controllers

import (
	"github.com/gin-gonic/gin"
	"sweetcook-backend/services"
	"sweetcook-backend/utils/error"
	"strconv"
	"sweetcook-backend/utils/logger"
)

type (
	Cookbook struct {
	}
)

/**
	获取菜谱列表功能
	
	http://localhost:7000/app-api/cookbook/list?page=0&num=10
 */
func (cb Cookbook)List(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	page,err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "参数错误。"})
		return
	}
	num,err := strconv.Atoi(c.Query("num"))
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "参数错误。"})
		return
	}
	cookbookList,err := services.QueryCookbookList(c, page, num)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "无更多数据。"})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.DATA_NAME: cookbookList})
}

/**
	添加菜谱到个人喜欢列表
	http://localhost:7000/app-api/cookbook/add
	{"ids":["1","10003"]}
 */
func (cb Cookbook)AddCookbookList(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	cookIds := services.UserCookbook{}
	if err := c.BindJSON(&cookIds);err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "参数错误。"})
		return
	}
	logger.Debug("ids is: ", cookIds)
	
	cookIds.Username = sessionUsername
	retMsg, err := services.AddCookbookList(c, cookIds)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: retMsg})
}

/**
	获取个人菜谱列表
	http://localhost:7000/app-api/cookbook/user_cookbooks
 */
func (cb Cookbook)QueryUserCookbooks(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	userCookbooks, err := services.QueryUserCookbooks(c, sessionUsername)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.DATA_NAME: userCookbooks})
}

/**
	获取情侣菜谱列表
	http://localhost:7000/app-api/cookbook/companion_cookbooks
 */
func (cb Cookbook)QueryCompanionCookbooks(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	//查询用户信息获取伴侣
	sessionUser,err := services.QueryOne(c, sessionUsername)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	
	if sessionUser.Companion == "" {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_NO_DATA_ERROR,error.MSG_NAME: "您尚未绑定伴侣。"})
		return
	}
	
	userCookbooks, err := services.QueryUserCookbooks(c, sessionUser.Companion)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.DATA_NAME: userCookbooks})
}