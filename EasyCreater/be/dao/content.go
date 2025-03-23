package dao

import (
	"demo/model"
	"log"
)

func AddContent(content *model.Content) error {
	res := db.Create(content)
	if res.Error != nil {
		log.Printf("Failed to add content: %v\n", res.Error)
		return res.Error
	}
	return nil
}

func GetContentByUserID(userID uint) ([]*model.Content, error) {
	var content []*model.Content
	res := db.Where("user_id = ?", userID).Order("timestamp desc").Find(&content)
	if res.Error != nil {
		log.Printf("Failed to get content: %v\n", res.Error)
		return nil, res.Error
	}
	return content, nil
}
