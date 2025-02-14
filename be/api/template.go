package api

import (
	"demo/dao"
	"demo/model"
	"demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SaveTemplate(c *gin.Context) {
	var temp model.ResumeTemplate
	if err := c.ShouldBindJSON(&temp); err != nil {
		utils.RespFail(c, "Verification failed!")
		return
	}
	name := temp.Name
	data := model.Data{
		temp.DataTemplate,
		temp.DataStyle,
		temp.DataScript,
	}

	err := dao.AddTemplate(name, data)
	if err != nil {
		utils.RespFail(c, "Add template failed!")
		return
	}
	utils.RespSuccess(c, "Add template success!")
	return
}

func GetTemplateFromId(c *gin.Context) {

	templateIdStr := c.Param("template_id")
	templateId, err := strconv.Atoi(templateIdStr)
	if err != nil {
		utils.RespFail(c, "template_id not found in context")
		return
	}

	template, err := dao.SearchTemplate(templateId)
	if err != nil {
		utils.RespFail(c, "search template failed!")
		return
	}

	data := model.Data{
		template.DataTemplate,
		template.DataStyle,
		template.DataScript,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":        200,
		"tempalte_name": template.Name,
		"data":          data,
	})
	return

}
