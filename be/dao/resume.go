package dao

import (
	"demo/model"
	"gorm.io/gorm"
	"log"
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
