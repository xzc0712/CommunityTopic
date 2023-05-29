package controller

import (
	"CommunityTopic_demo/service"
	"strconv"
)

type PageInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func QueryPageInfoById(id string) *PageInfo {
	topicId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return &PageInfo{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	pageInfo, err := service.GetPageData(int(topicId))
	if err != nil {
		return &PageInfo{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return &PageInfo{
		Code: 0,
		Msg:  "success",
		Data: pageInfo,
	}
}
