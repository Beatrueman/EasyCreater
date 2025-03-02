package api

import (
	"bytes"
	"demo/dao"
	"demo/model"
	"demo/utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"strconv"
	"strings"
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
		thumbnaiURL, err = uploadBase64ToOSS(requestData.Thumbnail, fmt.Sprintf("resume_%d_thumbnail.jpg", resumeId))
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

func uploadBase64ToOSS(base64Str, filename string) (string, error) {
	dao.LoadConfig()

	ossEndpoint := viper.GetString("OSS.Endpoint")
	ossAccessKey := viper.GetString("OSS.AccessKey")
	ossSecretKey := viper.GetString("OSS.SecretKey")
	ossBucketName := viper.GetString("OSS.BucketName")

	// 🎯 解析 Base64 数据
	dataIndex := strings.Index(base64Str, "base64,")
	if dataIndex == -1 {
		return "", fmt.Errorf("invalid base64 data")
	}
	base64Data := base64Str[dataIndex+7:] // 去掉前缀 "data:image/jpeg;base64,"

	decoded, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		return "", err
	}

	// 🎯 连接 OSS
	client, err := oss.New(ossEndpoint, ossAccessKey, ossSecretKey)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(ossBucketName)
	if err != nil {
		return "", err
	}

	// 🎯 上传到 OSS
	objectKey := "thumbnails/" + filename
	err = bucket.PutObject(objectKey, bytes.NewReader(decoded))
	if err != nil {
		return "", err
	}

	// 🎯 返回文件 URL
	url := fmt.Sprintf("https://yiiong.oss-cn-beijing.aliyuncs.com/%s", objectKey)
	return url, nil
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
