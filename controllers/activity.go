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
	
	ActivityComment struct {
		ActId        string `json:"act_id" bson:"act_id" binding:"required"`
		Rating    float32 `json:"rating" bson:"rating" binding:"required"`
		Comment    string   `json:"comment" bson:"comment" binding:"required"`
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
	
	http://localhost:7000/app-api/activity/delete?id=5950e12bf141511234bb1ae1
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

/**
	为煮饭活动评分接口
	
	http://localhost:7000/app-api/activity/comment_rate
	{"act_id":"5950ff0cf14151165ac68069","rating":8.4,"comment":"好好吃呀！"}
 */
func (act Activity)CommentAndRateActivity(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	activityComment := ActivityComment{}
	if err := c.BindJSON(&activityComment);err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_PARAM_ERROR,error.MSG_NAME: "参数错误。"})
		return
	}
	
	logger.Debug("activityComment is: ", activityComment)
	
	retMsg, err := services.AddActivityCommentRate(c, sessionUsername, activityComment.ActId, activityComment.Comment, activityComment.Rating)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: retMsg})
}

/**
	煮饭活动设置完成接口
	
	http://localhost:7000/app-api/activity/set_finished?act_id=5950ff0cf14151165ac68069
 */
func (act Activity)SetActivityFinished(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	actId := c.Query("act_id")
	retMsg, err := services.SetActivityFinished(c, sessionUsername, actId)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: retMsg})
}

/**
	煮饭活动单身狗访问接口
	
	http://localhost:7000/app-api/activity/single_dog?act_id=5950ff0cf14151165ac68069
 */
func (act Activity)SingleDogActivityInfo(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	actId := c.Query("act_id")
	retMsg, err := services.SingleDogActivityInfo(c, actId)
	if err != nil {
		c.JSON(200, gin.H{error.CODE_NAME: error.ERROR_SERVICE_ERROR,error.MSG_NAME: err.Error()})
		return
	}
	c.JSON(200, gin.H{error.CODE_NAME: error.SUCCESS,error.MSG_NAME: retMsg})
}


