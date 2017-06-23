package controllers

import (
	"github.com/gin-gonic/gin"
	"sweetcook-backend/services"
	"sweetcook-backend/utils/error"
	"strconv"
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

func (cb Cookbook)AddCookbookList(c *gin.Context)  {
	//拿出session
	sessionUsername := GetUserInfo(c)
	if sessionUsername == "" {
		return
	}
	
	
}