package dao

import (
	"demo/model"
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

//TODO: 编辑项
