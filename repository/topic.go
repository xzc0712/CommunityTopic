package repository

import "sync"

type Topic struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
}

type TopicDao struct {
}

var topicDao *TopicDao
var topicOnce sync.Once

func TopicDaoGet() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}

func (*TopicDao) QueryTopicById(id int) *Topic {
	return topicIndexMap[id]
}
