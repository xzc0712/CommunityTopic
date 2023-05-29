package service

import (
	"CommunityTopic_demo/repository"
	"errors"
	"log"
	"sync"
)

type PageData struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

type PageDataFlow struct {
	topicId  int
	pageData *PageData
	topic    *repository.Topic
	posts    []*repository.Post
}

func GetPageData(topicid int) (*PageData, error) {
	return NewGetPageDataFlow(topicid).Do()
}

func NewGetPageDataFlow(topicid int) *PageDataFlow {
	return &PageDataFlow{
		topicId: topicid,
	}
}

func (f *PageDataFlow) Do() (*PageData, error) {
	//做判断
	err := f.checkParam()
	if err != nil {
		log.Fatal(err)
	}
	//准备数据
	err = f.prepareData()
	if err != nil {
		log.Fatal(err)
	}
	//数据打包
	err = f.packPageData()
	if err != nil {
		log.Fatal(err)
	}
	return f.pageData, nil
}

func (f *PageDataFlow) checkParam() error {
	if f.topicId < 0 {
		return errors.New("topicId must larger than zero")
	}
	return nil
}

func (f *PageDataFlow) prepareData() error {
	var wg sync.WaitGroup
	wg.Add(2)
	//获取topic
	go func() {
		defer wg.Done()
		topic := repository.TopicDaoGet().QueryTopicById(f.topicId)
		f.topic = topic
	}()

	//获取post
	go func() {
		defer wg.Done()
		posts := repository.PostDaoGet().QueryPostById(f.topicId)
		f.posts = posts
	}()
	wg.Wait()
	return nil
}

func (f *PageDataFlow) packPageData() error {
	f.pageData = &PageData{
		Topic:    f.topic,
		PostList: f.posts,
	}
	return nil
}
