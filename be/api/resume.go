package api

import (
	"demo/dao"
	"demo/model"
	"demo/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func AddResume(c *gin.Context) {
	var requestData struct {
		ResumeData string `json:"resume_data"`
		Thumbnail  string `json:"thumbnail"`
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

	// 直接解析为 map[string]interface{}
	var resumeJSON map[string]interface{}
	if err := json.Unmarshal([]byte(requestData.ResumeData), &resumeJSON); err != nil {
		utils.RespFail(c, "Invalid resume JSON")
		return
	}

	templateName, ok := resumeJSON["templateName"].(string)
	if !ok {
		log.Printf("Invalid resume data: %v\n", resumeJSON)
		utils.RespFail(c, "Invalid resume data")
		return
	}

	ResumeName, ok := resumeJSON["resumeName"].(string)
	if !ok {
		log.Printf("Invalid resume data: %v\n", resumeJSON)
		utils.RespFail(c, "Invalid resume data")
		return
	}

	resume := model.ResumeData{
		Username:     username.(string),
		UserID:       userId,
		TemplateName: templateName,
		Resume:       requestData.ResumeData,
		IsShared:     false,
		ResumeName:   ResumeName,
	}

	if err := dao.AddResumeData(resume); err != nil {
		utils.RespFail(c, "Failed to add resume data")
		return
	}

	CurrentResume, err := dao.GetResumeDataLatest()
	if err != nil {
		utils.RespFail(c, "Failed to retrieve current resume data")
		return
	}

	resumeId := CurrentResume[len(CurrentResume)-1].ResumeId

	var thumbnaiURL string
	if requestData.Thumbnail != "" {
		thumbnaiURL, err = dao.UploadBase64ToOSS(requestData.Thumbnail, fmt.Sprintf("resume_%d_thumbnail.jpg", resumeId))
		log.Println(thumbnaiURL)
		if err != nil {
			utils.RespFail(c, "Failed to upload thumbnail")
			return
		}

		if err := dao.UpdateResumeThumbnail(resumeId, thumbnaiURL); err != nil {
			utils.RespFail(c, "Failed to update resume thumbnail")
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Resume data added successfully!",
		"data":   resumeId,
		"url":    thumbnaiURL,
	})
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

func GetThumbnail(c *gin.Context) {
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

	url, err := dao.GetResumeThunailURL(resumeId)
	if err != nil {
		utils.RespFail(c, "Failed to retrieve resume data")
		return
	}

	if url == "" {
		utils.RespFail(c, "No thumbnail found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   url,
	})
}

func uploadResume(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		utils.RespFail(c, "Failed to get file")
		return
	}
	defer file.Close()

	fileName := header.Filename
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}

	userIdStr, err := dao.SelectUserInfo(username.(string), "Id")
	if err != nil {
		utils.RespFail(c, "Failed to get user id")
		return
	}

	userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		utils.RespFail(c, "Failed to parse user id")
		return
	}

	ossURL, err := dao.UploadToOSS(file, fileName, username.(string))
	if err != nil {
		utils.RespFail(c, "Failed to upload file")
		return
	}

	loadedResume := model.LoadedResumeData{
		Username:   username.(string),
		URL:        ossURL,
		UserID:     int(userIdInt),
		ResumeName: fileName,
		Timestamp:  time.Now(),
	}

	err = dao.AddLoadedResumeData(loadedResume)
	if err != nil {
		utils.RespFail(c, "Failed to add loaded resume data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "File uploaded successfully!",
		"data": gin.H{
			"url": ossURL,
		},
	})
}

func GetLoadedResume(c *gin.Context) {
	username, ok := c.Get("username")
	if !ok {
		utils.RespFail(c, "Username not found in context")
		return
	}

	loadedResumeData, err := dao.GetLoadedResumeData(username.(string))
	if err != nil {
		utils.RespFail(c, "Failed to retrieve loaded resume data")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "File uploaded successfully!",
		"data":   loadedResumeData,
	})
}

func DeleteLoadedResume(c *gin.Context) {
	resumeIdStr := c.Param("resume_id")
	if resumeIdStr == "" {
		utils.RespFail(c, "Resume ID is required")
		return
	}

	resumeId, err := strconv.Atoi(resumeIdStr)
	if err != nil {
		utils.RespFail(c, "Invalid Resume ID")
		return
	}

	err = dao.DeleteLoadedResumeData(resumeId)
	if err != nil {
		utils.RespFail(c, "Failed to delete loaded resume data")
		return
	}
	utils.RespSuccess(c, "Loaded resume data deleted successfully!")
}

func GetLoadedResumeURL(c *gin.Context) {
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

	url, err := dao.GetLoadedResumeURL(resumeId)
	if err != nil {
		utils.RespFail(c, "Failed to retrieve resume data")
		return
	}

	if url == "" {
		utils.RespFail(c, "No url found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   url,
	})
}

func GetIdeaData(c *gin.Context) {
	ideaData, err := dao.GetIdea()
	if err != nil {
		utils.RespFail(c, "Failed to retrieve idea data")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   json.RawMessage(ideaData),
	})
}
