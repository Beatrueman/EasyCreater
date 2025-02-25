package api

import (
	"demo/dao"
	"demo/model"
	"demo/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddResume(c *gin.Context) {
	var requestData struct {
		ResumeData string `json:"resume_data"`
	}

	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.RespFail(c, "Invalid JSON")
		return
	}

	// 查询用户ID
	userIdstr, err := dao.SelectUserInfo(username.(string), "Id")
	if err != nil {
		utils.RespFail(c, "Error retrieving user ID")
		return
	}

	userId, err := strconv.Atoi(userIdstr)
	if err != nil {
		utils.RespFail(c, "Invalid user ID")
		return
	}

	// 解析resume_data，提取templateName
	var resumeJSON map[string]interface{}
	if err := json.Unmarshal([]byte(requestData.ResumeData), &resumeJSON); err != nil {
		utils.RespFail(c, "Invalid resume data")
		return
	}

	templateName, ok := resumeJSON["templateName"].(string)
	if !ok {
		utils.RespFail(c, "Invalid resume data")
		return
	}

	resume := model.ResumeData{
		Username:     username.(string),
		UserID:       userId,
		TemplateName: templateName,
		Resume:       requestData.ResumeData,
		IsShared:     false,
	}

	if err := dao.AddResumeData(resume); err != nil {
		utils.RespFail(c, "Failed to add resume data")
		return
	}

	utils.RespSuccess(c, "Resume data saved successfully!")
}

func GetResume(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}

	resumeData, err := dao.GetResumeData(username.(string))
	if err != nil {
		utils.RespFail(c, "Failed to retrieve resume data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   resumeData,
	})
}

func DeleteResume(c *gin.Context) {
	resumeIdstr := c.Param("resume_id")
	if resumeIdstr == "" {
		utils.RespFail(c, "Resume ID is required")
		return
	}

	resumeId, err := strconv.Atoi(resumeIdstr)
	if err != nil {
		utils.RespFail(c, "Invalid Resume ID")
		return
	}

	if err := dao.DeleteResumeData(resumeId); err != nil {
		utils.RespFail(c, "Failed to delete resume data")
		return
	}

	utils.RespSuccess(c, "Resume data deleted successfully!")
}

func GetResumeFromId(c *gin.Context) {
	resumeIdstr := c.Param("resume_id")
	if resumeIdstr == "" {
		utils.RespFail(c, "Resume ID is required")
		return
	}

	resumeId, err := strconv.Atoi(resumeIdstr)
	if err != nil {
		utils.RespFail(c, "Invalid Resume ID")
		return
	}
	resumeData, err := dao.GetResumeFromId(resumeId)
	if err != nil {
		utils.RespFail(c, "Failed to retrieve resume data")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   resumeData,
	})
}

func ShareResume(c *gin.Context) {
	resumeIdstr := c.Param("resume_id")
	if resumeIdstr == "" {
		utils.RespFail(c, "Resume ID is required")
		return
	}

	resumeId, err := strconv.Atoi(resumeIdstr)
	if err != nil {
		utils.RespFail(c, "Invalid Resume ID")
		return
	}

	var requestBody struct {
		Action string `json:"action"` // share or unshare
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.RespFail(c, "Invalid JSON")
		return
	}

	var isShared bool
	if requestBody.Action == "share" {
		isShared = true
	} else if requestBody.Action == "unshare" {
		isShared = false
	} else {
		utils.RespFail(c, "Invalid action")
		return
	}

	if err := dao.ShareResumeUpdate(resumeId, isShared); err != nil {
		utils.RespFail(c, "Failed to share resume data")
		return
	}

	utils.RespSuccess(c, "Resume data shared successfully!")
}

func GetSharedResume(c *gin.Context) {
	isAll := c.DefaultQuery("is_all", "false")
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}
	var resumeData []*model.ResumeData
	var err error

	// 根据 is_all 的值来判断查询条件
	if isAll == "true" {
		// 查询所有已分享的简历
		resumeData, err = dao.GetSharedResume("")
	} else {
		// 查询当前用户的已分享简历
		resumeData, err = dao.GetSharedResume(username.(string))
	}

	if err != nil {
		utils.RespFail(c, "Failed to retrieve shared resume data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   resumeData,
	})
}
