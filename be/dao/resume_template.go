package dao

import (
	"demo/model"
	"log"
)

// 添加简历模板

func AddTemplate(name string, data model.Data) error {

	template := model.ResumeTemplate{
		Name:         name,
		DataTemplate: data.DataTemplate,
		DataStyle:    data.DataStyle,
		DataScript:   data.DataScript,
	}

	res := db.Create(&template)
	if res.Error != nil {
		log.Fatalf("Error saving template: %s\n", res.Error)
		return res.Error
	}

	return nil
}

// 根据 Id 查找简历模板

func SearchTemplate(templateId int) (model.ResumeTemplate, error) {
	var template model.ResumeTemplate
	res := db.Where("template_id = ?", templateId).First(&template)
	if res.Error != nil {
		log.Fatalf("Error searching template: %s\n", res.Error)
		return template, res.Error
	}
	return template, nil
}
