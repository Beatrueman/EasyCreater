package dao

import (
	"demo/model"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"net/url"
	"strings"
)

func AddResumeData(resume model.ResumeData) error {
	res := db.Create(&resume)
	if res.Error != nil {
		log.Printf("Failed to add resume data: %v\n", res.Error)
	}
	return nil
}

func GetResumeData(username string) ([]*model.ResumeData, error) {
	var resumeData []*model.ResumeData
	res := db.Where("username = ?", username).Order("timestamp desc").Find(&resumeData)
	if res.Error != nil {
		log.Printf("Failed to get resume data: %v\n", res.Error)
		return nil, res.Error
	}
	return resumeData, nil
}

func GetResumeDataLatest() ([]*model.ResumeData, error) {
	var resumeData []*model.ResumeData
	res := db.Last(&resumeData)
	if res.Error != nil {
		log.Printf("Failed to get resume data: %v\n", res.Error)
		return nil, res.Error
	}
	return resumeData, nil
}

func GetResumeFromId(ResumeId int) ([]*model.ResumeData, error) {
	var resumeData []*model.ResumeData
	res := db.Where("resume_id = ?", ResumeId).Find(&resumeData)
	if res.Error != nil {
		log.Printf("Failed to get resume data: %v\n", res.Error)
		return nil, res.Error
	}
	return resumeData, nil
}

func DeleteResumeData(ResumeId int) error {
	var resumeData []*model.ResumeData
	res := db.Where("resume_id = ?", ResumeId).Delete(&resumeData)
	if res.Error != nil {
		log.Printf("Failed to delete resume data: %v\n", res.Error)
		return res.Error
	}
	return nil
}

// 分享简历
func ShareResumeUpdate(ResumeId int, isShared bool) error {
	res := db.Model(&model.ResumeData{}).Where("resume_id = ?", ResumeId).Update("is_shared", isShared)
	if res.Error != nil {
		log.Printf("Failed to update resume data: %v\n", res.Error)
		return res.Error
	}
	return nil
}

// 获取某个用户已分享的简历
func GetSharedResume(username string) ([]*model.ResumeData, error) {
	var resumeData []*model.ResumeData
	var res *gorm.DB

	if username != "" {
		res = db.Where("is_shared = ? AND username = ?", true, username).Find(&resumeData)
	} else {
		res = db.Where("is_shared = ? ", true).Find(&resumeData)
	}

	if res.Error != nil {
		log.Printf("Failed to get shared resume data: %v\n", res.Error)
		return nil, res.Error
	}
	return resumeData, nil
}

// 上传缩略图到oss
func UpdateResumeThumbnail(ResumeId int, ThumbnailUrl string) error {
	res := db.Model(&model.ResumeData{}).Where("resume_id = ?", ResumeId).Update("thumbnail_url", ThumbnailUrl)
	if res.Error != nil {
		log.Printf("Failed to update resume thumbnail: %v\n", res.Error)
		return res.Error
	}
	return nil
}

// 根据resume_id 查询 thumbnail_url
func GetResumeThunailURL(ResumeId int) (string, error) {
	// 读取oss配置
	LoadConfig()

	ossEndpoint := viper.GetString("OSS.Endpoint")
	ossAccessKey := viper.GetString("OSS.AccessKey")
	ossSecretKey := viper.GetString("OSS.SecretKey")
	ossBucketName := viper.GetString("OSS.BucketName")

	var thumbnailUrl string

	err := db.Table("resume_data").Where("resume_id = ?", ResumeId).Select("thumbnail_url").Limit(1).Pluck("thumbnail_url", &thumbnailUrl).Error
	if err != nil {
		log.Printf("Failed to get resume thumbnail: %v\n", err)
		return "", err
	}

	if thumbnailUrl == "" {
		log.Println("Thumbnail URL not found")
		return "", errors.New("thumbnail URL not found")
	}

	parsedURL, err := url.Parse(thumbnailUrl)
	if err != nil {
		return "", err
	}

	objectKey := strings.TrimPrefix(parsedURL.Path, "/")

	// 连接 OSS
	client, err := oss.New(ossEndpoint, ossAccessKey, ossSecretKey)
	if err != nil {
		log.Println("OSS 连接失败:", err)
		return "", err
	}

	bucket, err := client.Bucket(ossBucketName)
	if err != nil {
		log.Println("获取 OSS Bucket 失败:", err)
		return "", err
	}

	// 生成签名 URL（有效期 1 小时）
	signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 3600)
	if err != nil {
		log.Println("生成签名 URL 失败:", err)
		return "", err
	}

	return signedURL, nil
}
