package controller

import (
	"CommunityTopic_demo/service"
	"log"
	"strconv"
)

//存入新帖子

func PublishPost(topicIdStr string, content string) *PageInfo {
	topicId, err := strconv.ParseInt(topicIdStr, 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	//获取service层结果
	postId, err := service.PublishPost(int(topicId), content)
	if err != nil {
		return &PageInfo{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageInfo{
		Code: 1,
		Msg:  "success",
		Data: map[string]int{
			"post_id": postId,
		},
	}

}
