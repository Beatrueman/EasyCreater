package dao

import (
	"bytes"
	"demo/model"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"io"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	log.Printf("Sharing resume ID: %d, new status: %v", ResumeId, isShared)

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

// 根据 resume_id 查询 thumbnail_url
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

// 上传用户文件到 OSS
func UploadToOSS(fileData multipart.File, filename string, username string) (string, error) {
	LoadConfig()

	ossEndpoint := viper.GetString("OSS.Endpoint")
	ossAccessKey := viper.GetString("OSS.AccessKey")
	ossSecretKey := viper.GetString("OSS.SecretKey")
	ossBucketName := viper.GetString("OSS.BucketName")

	// 🎯 连接 OSS
	client, err := oss.New(ossEndpoint, ossAccessKey, ossSecretKey)
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(ossBucketName)
	if err != nil {
		return "", err
	}

	// 🎯 生成唯一文件名
	timestamp := time.Now().Unix()
	ext := filepath.Ext(filename) // 获取文件扩展名
	uniqueFilename := fmt.Sprintf("%d%s", timestamp, ext)

	// 🎯 上传到 OSS
	objectKey := fmt.Sprintf("userResumes/%s/%s", username, uniqueFilename)
	err = bucket.PutObject(objectKey, fileData)
	if err != nil {
		return "", err
	}

	// 🎯 返回文件 URL
	url := fmt.Sprintf("https://%s.%s/%s", ossBucketName, ossEndpoint, objectKey)
	return url, nil
}

func UploadBase64ToOSS(base64Str, filename string) (string, error) {
	LoadConfig()

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

func AddLoadedResumeData(resume model.LoadedResumeData) error {
	res := db.Create(&resume)
	if res.Error != nil {
		log.Printf("Failed to add resume data: %v\n", res.Error)
	}
	return nil
}

func GetLoadedResumeData(username string) ([]*model.LoadedResumeData, error) {
	var resumeData []*model.LoadedResumeData
	res := db.Where("username = ?", username).Find(&resumeData)
	if res.Error != nil {
		log.Printf("Failed to get resume data: %v\n", res.Error)
		return nil, res.Error
	}
	return resumeData, nil
}

func DeleteLoadedResumeData(resumeId int) error {
	res := db.Where("resume_id = ?", resumeId).Delete(&model.LoadedResumeData{})
	if res.Error != nil {
		log.Printf("Failed to delete resume data: %v\n", res.Error)
		return res.Error
	}
	return nil
}

func GetLoadedResumeURL(ResumeId int) (string, error) {
	// 读取oss配置
	LoadConfig()

	ossEndpoint := viper.GetString("OSS.Endpoint")
	ossAccessKey := viper.GetString("OSS.AccessKey")
	ossSecretKey := viper.GetString("OSS.SecretKey")
	ossBucketName := viper.GetString("OSS.BucketName")

	var Url string

	err := db.Table("loaded_resume_data").Where("resume_id = ?", ResumeId).Select("url").Limit(1).Pluck("url", &Url).Error
	if err != nil {
		log.Printf("Failed to get resume thumbnail: %v\n", err)
		return "", err
	}

	if Url == "" {
		log.Println("URL not found")
		return "", errors.New("URL not found")
	}

	parsedURL, err := url.Parse(Url)
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
	log.Println("签名 URL:", signedURL)

	return signedURL, nil
}

func GetIdea() ([]byte, error) {
	// 读取idea.json
	dir, _ := os.Getwd()
	fmt.Println("当前运行目录是:", dir)

	file, err := os.Open("idea.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

// ToggleResumeLike 切换简历点赞状态
func ToggleResumeLike(resumeID, userID int, username string) (bool, error) {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var existingLike model.ResumeLike
	err := tx.Where("resume_id = ? AND user_id = ?", resumeID, userID).First(&existingLike).Error
	if err == gorm.ErrRecordNotFound {
		// 点赞，插入记录
		newLike := model.ResumeLike{
			ResumeID: resumeID,
			UserID:   userID,
			Username: username,
		}
		if err := tx.Create(&newLike).Error; err != nil {
			tx.Rollback()
			return false, err
		}

		// 点赞计数 +1
		if err := tx.Model(&model.ResumeData{}).
			Where("resume_id = ?", resumeID).
			Update("like_count", gorm.Expr("like_count + 1")).Error; err != nil {
			tx.Rollback()
			return false, err
		}

		if err := tx.Commit().Error; err != nil {
			return false, err
		}
		return true, nil
	} else if err != nil {
		tx.Rollback()
		return false, err
	} else {
		// 取消点赞，删除记录
		if err := tx.Delete(&existingLike).Error; err != nil {
			tx.Rollback()
			return false, err
		}

		// 点赞计数减1，防止负数
		if err := tx.Model(&model.ResumeData{}).
			Where("resume_id = ?", resumeID).
			Update("like_count", gorm.Expr("CASE WHEN like_count > 0 THEN like_count - 1 ELSE 0 END")).Error; err != nil {
			tx.Rollback()
			return false, err
		}

		if err := tx.Commit().Error; err != nil {
			return false, err
		}
		return false, nil
	}
}

// GetResumeLikeStatus 获取用户对特定简历的点赞状态
func GetResumeLikeStatus(resumeID, userID int) (bool, error) {
	var like model.ResumeLike
	result := db.Where("resume_id = ? AND user_id = ?", resumeID, userID).First(&like)

	if result.Error == gorm.ErrRecordNotFound {
		return false, nil // 未点赞
	} else if result.Error != nil {
		log.Printf("Database error: %v\n", result.Error)
		return false, result.Error
	}

	return true, nil // 已点赞
}

// GetResumeLikeCount 获取简历点赞数量
func GetResumeLikeCount(resumeID int) (int, error) {
	var resume model.ResumeData
	result := db.Select("like_count").Where("resume_id = ?", resumeID).First(&resume)

	if result.Error != nil {
		log.Printf("Failed to get like count: %v\n", result.Error)
		return 0, result.Error
	}

	return resume.LikeCount, nil
}
