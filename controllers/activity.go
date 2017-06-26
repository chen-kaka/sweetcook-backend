package controllers

import (
	"github.com/gin-gonic/gin"
	"sweetcook-backend/services"
	"sweetcook-backend/utils/error"
	"sweetcook-backend/utils/logger"
	"strconv"
)

type (
	Activity struct {
	}
)

/**
	创建或修改煮饭活动接口
	
	http://localhost:7000/app-api/activity/create_update
	{"_id":"efsdfsd","cook_ids":["1001","1002"],"cook_user":"kaka","content":"好好煮饭~"}
 */
func (act Activity)CreateOrUpdateActivity(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	activity := services.Activity{}
	if err := c.BindJSON(&activity);err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "参数错误。"})
		return
	}
	logger.Debug("activity is: ", activity)
	
	activity.Username = sessionUsername
	retMsg, err := services.CreateOrUpdateActivity(c, activity)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: retMsg})
}

/**
	查看煮饭活动接口

	http://localhost:7000/app-api/activity/list?page=0&num=10
 */
func (act Activity)ListActivity(c *gin.Context)  {
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
	activityList,err := services.QueryActivityList(c, sessionUsername, page, num)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "无更多数据。"})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.DATA_NAME: activityList})
}

/**
	删除煮饭活动接口
	
	http://localhost:7000/app-api/activity/delete?id=sdfsdfsd
 */
func (act Activity)DeleteActivity(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	id := c.Query("id")
	
	if id == "" {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "参数错误。"})
		return
	}
	logger.Debug("id is: ", id)
	
	retMsg, err := services.DeleteActivity(c, sessionUsername, id)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: retMsg})
}