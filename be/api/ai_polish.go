package api

import (
	"context"
	"demo/dao"
	"demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	easyllm "github.com/soryetong/go-easy-llm"
	"github.com/soryetong/go-easy-llm/easyai/chatmodule"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"time"
)

type ResumeRequest struct {
	ResumeData string `json:"resume_data"`
}

type AIResponse struct {
	Reply     string `json:"reply"`
	RequestId string `json:"request_id"`
}

func CacheResume(c *gin.Context) {
	var req ResumeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("error binding JSON:", req)
		utils.RespFail(c, "Invalid JSON")
		return
	}
	taskID := uuid.New().String()
	if err := dao.SetResumeCache(taskID, req.ResumeData, 10*time.Minute); err != nil {
		utils.RespFail(c, "Failed to cache resume's taskID")
		return
	}
	c.JSON(http.StatusOK,
		gin.H{
			"data": gin.H{
				"task_id": taskID,
			},
			"status": 200,
		})

}
func QWenStreamChat(c *gin.Context) {

	dao.LoadConfig()

	token := viper.GetString("token")
	taskID := c.Query("task_id")
	if taskID == "" {
		c.Header("Content-Type", "text/event-stream")
		c.Writer.Write([]byte("event: error\n"))
		c.Writer.Write([]byte("data: Missing task_id\n\n"))
		c.Writer.Flush()
		return
	}

	resumeData, err := dao.GetResumeCache(taskID)
	if err != nil {
		c.Header("Content-Type", "text/event-stream")
		c.Writer.Write([]byte("event: error\n"))
		c.Writer.Write([]byte("data: Failed to get resume's taskID\n\n"))
		c.Writer.Flush()
		return
	}

	globalParams := new(chatmodule.QWenParameters)
	globalParams.Input = &chatmodule.QWenInputMessages{}
	tipsMsg := &chatmodule.ChatMessage{Role: chatmodule.IdSystem, Content: "你是一个简历优化助手，你需要为我优化json格式的简历内容，为了保证体验感，请不要输出一切json格式的内容，不要输出优化后的简历内容，不要用markdown语法。提出的建议需要详细"}
	globalParams.Input.Messages = append(globalParams.Input.Messages, tipsMsg)
	globalParams.Parameters = map[string]interface{}{
		"temperature": 0.8,
		"top_p":       0.8,
		"max_tokens":  1500,
	}

	config := easyllm.DefaultConfig(token, chatmodule.ChatTypeQWen)
	client := easyllm.NewChatClient(config).SetGlobalParams(globalParams)
	stream, err := client.StreamChat(context.Background(), &chatmodule.ChatRequest{
		Model:   "qwen-plus-2025-01-25",
		Message: resumeData,
	})
	if err != nil {
		log.Println("Error sending chat request:", err)
		c.Header("Content-Type", "text/event-stream")
		c.Writer.Write([]byte("event: error\n"))
		c.Writer.Write([]byte("data: Error sending chat request\n\n"))
		c.Writer.Flush()
		return
	}

	// 设置 SSE 头部
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Writer.Flush()

	// 启动流式输出
	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-stream:
			if !ok {
				// channel 关闭了，结束流
				return false
			}
			if msg == nil {
				return true // 继续等待下一条消息
			}
			c.SSEvent("", gin.H{
				"content":    msg.Content,
				"session_id": msg.SessionId,
			})
			return true // 继续监听
		case <-c.Request.Context().Done():
			// 客户端断开连接，停止推送
			return false
		}
	})

}

func QWenStreamChatBase(c *gin.Context) {

	dao.LoadConfig()

	token := viper.GetString("token")

	taskID := c.Query("task_id")
	if taskID == "" {
		c.Header("Content-Type", "text/event-stream")
		c.Writer.Write([]byte("event: error\n"))
		c.Writer.Write([]byte("data: Missing task_id\n\n"))
		c.Writer.Flush()
		return
	}

	resumeData, err := dao.GetResumeCache(taskID)
	if err != nil {
		c.Header("Content-Type", "text/event-stream")
		c.Writer.Write([]byte("event: error\n"))
		c.Writer.Write([]byte("data: Failed to get resume's taskID\n\n"))
		c.Writer.Flush()
		return
	}

	globalParams := new(chatmodule.QWenParameters)
	globalParams.Input = &chatmodule.QWenInputMessages{}
	tipsMsg := &chatmodule.ChatMessage{Role: chatmodule.IdSystem, Content: "你是一个简历助手，你需要提供一些写简历的建议，以帮助我可以做出一份不错的简历。" +
		"为了保证体验感，请不要输出一切json格式的内容。提出的建议需要详细一些"}
	globalParams.Input.Messages = append(globalParams.Input.Messages, tipsMsg)
	globalParams.Parameters = map[string]interface{}{
		"temperature": 0.8,
		"top_p":       0.8,
		"max_tokens":  1500,
	}

	Question := resumeData
	if Question == "" {
		Question = "我是一个准备求职的大学生，现在需要写简历，但是我不会写，你能给我提出一些建议吗?"
	} else {
		Question = Question + "这是我目前的个人信息，请帮我生成一份简历建议。"
	}

	log.Printf(Question)
	config := easyllm.DefaultConfig(token, chatmodule.ChatTypeQWen)
	client := easyllm.NewChatClient(config).SetGlobalParams(globalParams)
	stream, err := client.StreamChat(context.Background(), &chatmodule.ChatRequest{
		Model:   "qwen-plus-2025-01-25",
		Message: Question,
	})
	if err != nil {
		log.Println("Error sending chat request:", err)
		c.Header("Content-Type", "text/event-stream")
		c.Writer.Write([]byte("event: error\n"))
		c.Writer.Write([]byte("data: Error sending chat request\n\n"))
		c.Writer.Flush()
		return
	}

	// 设置 SSE 头部
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Writer.Flush()

	// 启动流式输出
	c.Stream(func(w io.Writer) bool {
		select {
		case msg, ok := <-stream:
			if !ok {
				// channel 关闭了，结束流
				return false
			}
			if msg == nil {
				return true // 继续等待下一条消息
			}
			c.SSEvent("", gin.H{
				"content":    msg.Content,
				"session_id": msg.SessionId,
			})
			return true // 继续监听
		case <-c.Request.Context().Done():
			// 客户端断开连接，停止推送
			return false
		}
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
