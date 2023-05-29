package main

import (
	"CommunityTopic_demo/controller"
	"CommunityTopic_demo/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := Init("data/")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()

	//获取post和topic信息
	//127.0.0.1:8080/pageGet/2
	r.GET("pageGet/:id", func(c *gin.Context) {
		id := c.Param("id")
		data := controller.QueryPageInfoById(id)
		c.JSON(200, data)
	})

	//新增发帖内容,传入主体id和该帖内容
	//127.0.0.1:8080/pagePost/submit
	r.POST("/pagePost/submit", func(c *gin.Context) {
		topicId := c.PostForm("topicId")
		postContent := c.PostForm("content")
		data := controller.PublishPost(topicId, postContent)
		c.JSON(200, data)
	})
	err = r.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// Init 初始化数据内容
func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
