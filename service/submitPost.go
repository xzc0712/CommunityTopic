package service

import (
	"CommunityTopic_demo/repository"
	"errors"
	idworker "github.com/gitstliu/go-id-worker"
	"time"
)

type PublishNewPostFlow struct {
	content string
	topicId int
	postId  int
}

func PublishPost(topicId int, content string) (int, error) {
	return NewPublishPost(topicId, content).Do()
}

func NewPublishPost(topicId int, content string) *PublishNewPostFlow {
	return &PublishNewPostFlow{
		topicId: topicId,
		content: content,
	}
}

func (f PublishNewPostFlow) Do() (int, error) {
	if err := f.checkParam(); err != nil {
		return -1, err
	}
	if err := f.publish(); err != nil {
		return -1, err
	}
	return f.postId, nil
}

//检查内容长度是否符合要求
func (f PublishNewPostFlow) checkParam() error {
	if len(f.content) >= 500 {
		return errors.New("content length must be less than 500")
	}
	return nil
}

var idGen *idworker.IdWorker

//ID生成器的初始化
func init() {
	idGen = &idworker.IdWorker{}
	idGen.InitIdWorker(1, 1) //WORKERID位数 (用于对工作进程进行编码), 数据中心ID位数 (用于对数据中心进行编码)
}

//调用Dao层，将新帖存入数据库
func (f PublishNewPostFlow) publish() error {
	post := &repository.Post{
		ParentId:   f.topicId,
		Content:    f.content,
		CreateTime: int(time.Now().Unix()),
	}
	id, err := idGen.NextId()
	if err != nil {
		return err
	}
	post.Id = int(id)
	//调用Dao层
	err = repository.PostDaoPost().InsertNewPost(post)
	if err != nil {
		return err
	}
	f.postId = post.Id
	return nil

}
