package api

import (
	"context"
	"demo/utils"
	"github.com/gin-gonic/gin"
	easyllm "github.com/soryetong/go-easy-llm"
	"github.com/soryetong/go-easy-llm/easyai"
	"log"
	"net/http"
)

type ResumeRequest struct {
	ResumeData string `json:"resume_data"`
}

type AIResponse struct {
	Reply     string `json:"reply"`
	RequestId string `json:"request_id"`
}

func QWenNormalChat(c *gin.Context) {

	var req ResumeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error binding JSON:", req)
		utils.RespFail(c, "Invalid JSON")
		return
	}

	globalParams := new(easyai.QWenParameters)
	globalParams.Input = &easyai.QWenInputMessages{}
	tipsMsg := &easyai.ChatMessage{Role: easyai.IdSystem, Content: "你是一个简历优化助手，你需要为我优化json格式的简历内容，为了保证体验感，请不要输出一切json格式的内容，不要输出优化后的简历内容，不要用markdown语法。提出的建议需要详细"}
	globalParams.Input.Messages = append(globalParams.Input.Messages, tipsMsg)
	globalParams.Parameters = map[string]interface{}{
		"temperature": 0.8,
		"top_p":       0.8,
		"max_tokens":  1500,
	}

	config := easyllm.DefaultConfig("sk-70afc0771fc44f01a42ed73e983a6547", easyai.ChatTypeQWen)
	client := easyllm.NewChatClient(config).SetGlobalParams(globalParams)
	resp, reply, err := client.NormalChat(context.Background(), &easyai.ChatRequest{
		Model:   easyai.ChatModelQWenTurbo,
		Message: req.ResumeData,
	})
	if err != nil {
		log.Println("Error sending chat request:", err)
		utils.RespFail(c, "Error sending chat request")
		return
	}

	// 返回建议
	c.JSON(http.StatusOK, AIResponse{
		Reply:     resp.Content,
		RequestId: reply.(*easyai.QWenResponse).RequestId,
	})

}

func QWenNormalChatBase(c *gin.Context) {

	var req ResumeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespFail(c, "Invalid JSON")
		return
	}

	globalParams := new(easyai.QWenParameters)
	globalParams.Input = &easyai.QWenInputMessages{}
	tipsMsg := &easyai.ChatMessage{Role: easyai.IdSystem, Content: "你是一个简历助手，你需要提供一些写简历的建议，以帮助我可以做出一份不错的简历。" +
		"为了保证体验感，请不要输出一切json格式的内容，不要用markdown语法。提出的建议需要详细一些"}
	globalParams.Input.Messages = append(globalParams.Input.Messages, tipsMsg)
	globalParams.Parameters = map[string]interface{}{
		"temperature": 0.8,
		"top_p":       0.8,
		"max_tokens":  1500,
	}

	Question := req.ResumeData
	if Question == "" {
		Question = "我是一个准备求职的大学生，现在需要写简历，但是我不会写，你能给我提出一些建议吗?"
	} else {
		Question = Question + "这是我目前的个人信息，请帮我生成一份简历建议。"
	}

	log.Printf(Question)
	config := easyllm.DefaultConfig("sk-70afc0771fc44f01a42ed73e983a6547", easyai.ChatTypeQWen)
	client := easyllm.NewChatClient(config).SetGlobalParams(globalParams)
	resp, reply, err := client.NormalChat(context.Background(), &easyai.ChatRequest{
		Model:   easyai.ChatModelQWenTurbo,
		Message: Question,
	})
	if err != nil {
		log.Println("Error sending chat request:", err)
		utils.RespFail(c, "Error sending chat request")
		return
	}

	//fmt.Printf("resp: %+v\n", resp)

	// 返回建议
	c.JSON(http.StatusOK, AIResponse{
		Reply:     resp.Content,
		RequestId: reply.(*easyai.QWenResponse).RequestId,
	})

}

//func TestQWenStreamChat(t *testing.T) {
//	globalParams := new(easyai.QWenParameters)
//	globalParams.Input = &easyai.QWenInputMessages{}
//	tipsMsg := &easyai.ChatMessage{Role: easyai.IdSystem, Content: "你是一个简历优化助手，你需要为我优化json格式的简历内容，为了保证体验感，请尽量不要输出json结构体，提出的建议需要详细"}
//	globalParams.Input.Messages = append(globalParams.Input.Messages, tipsMsg)
//
//	config := easyllm.DefaultConfig("sk-70afc0771fc44f01a42ed73e983a6547", easyai.ChatTypeQWen)
//	client := easyllm.NewChatClient(config)
//	client.SetCustomParams(globalParams)
//	resp, err := client.StreamChat(context.Background(), &easyai.ChatRequest{
//		Model:   easyai.ChatModelQWenTurbo,
//		Message: "\"colors\":{\"left\":{\"highlight\":\"#82C0CC\",\"text\":\"white\",\"background\":\"#3943B7\"},\"right\":{\"highlight\":\"#3943B7\",\"text\":\"black\",\"background\":\"white\"}},\"name\":\"Yiiong\",\"title\":\"SRE工程师\",\"introText\":\"个人介绍\",\"imgUrl\":\"/avatar.png\",\"headlines\":[\"个人信息\",\"联系方式\",\"专业技能\",\"荣誉/证书\",\"教育经历\",\"项目经历\",\"工作经历\"],\"honor\":[\"国家一等奖学金\",\"HCIE认证\"],\"contact\":{\"phone\":\"19823838711\",\"email\":\"yiiong@qq.com\",\"address\":\"重庆\"},\"skills\":[\"Python\",\"Linux\",\"Git\",\"Docker\",\"Kubernetes\",\"Go\"],\"education\":[{\"title\":\"本科\",\"university\":\"重庆邮电大学\",\"major\":\"计算机科学与技术\",\"date\":\"2022-2026\",\"description\":\"主修课程：计算机网络、数据结构、操作系统、计算机组成原理\"},{\"title\":\"硕士\",\"university\":\"重庆邮电大学\",\"major\":\"计算机科学与技术\",\"date\":\"2022-2026\",\"description\":\"从事的研究方向\"}],\"experience\":[{\"company\":\"ABC公司\",\"position\":\"算法工程师\",\"location\":\"重庆\",\"date\":\"2022-至今\",\"description\":[\"1. 搭建云服务器，使用Python进行开发\",\"2. 使用Python进行机器学习\",\"3. 使用Python进行机器学习\"]},{\"company\":\"XXX公司\",\"position\":\"云计算工程师\",\"location\":\"重庆\",\"date\":\"2022-至今\",\"description\":[\"1. 搭建云服务器，使用Python进行开发\",\"2. Kubernetes 部署\"]}],\"project_experience\":[{\"name\":\"在线简历生成器\",\"tech_stack\":\"Python、机器学习、Kubernetes、Docker、Linux、Git\",\"description\":[\"1. 使用Vue 开发\",\"2. 后端使用Go 开发\",\"3. 在 Kubernetes 部署\"]}],\"editing\":false,\"showImage\":true,\"widthLeft\":30,\"imageShape\":\"circle\",\"headlineWeight\":\"400\",\"resumeFormat\":\"a4\"}",
//	})
//	if err != nil {
//		t.Log(err)
//		return
//	}
//
//	markdownFilterSrv := new(service.MarkdownProcessor)
//	for content := range resp {
//		t.Log("content: ", content)
//
//		if markdownFilterSrv.Do(content.Content) != "" {
//			t.Log("content.Content", content.Content)
//		}
//	}
//}
