package repository

import (
	"encoding/json"
	"os"
	"sync"
)

type Post struct {
	Id         int    `json:"id"`
	ParentId   int    `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
}

type PostDao struct {
}

var postDao *PostDao
var postOnce sync.Once
var rwMutex sync.RWMutex

func PostDaoGet() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func PostDaoPost() *PostDao {
	postOnce.Do(
		func() {
			postDao = &PostDao{}
		})
	return postDao
}

func (*PostDao) QueryPostById(id int) []*Post {
	return postsIndexMap[id]
}

func (*PostDao) InsertNewPost(post *Post) error {
	open, err := os.OpenFile("./data/post", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer open.Close()
	postStr, _ := json.Marshal(post) //转换成json格式

	//写文件时需要string
	//注意windows中回车换行符
	if _, err = open.WriteString(string(postStr) + "\n"); err != nil {
		return err
	}
	//更新Map时考虑并发安全性问题
	rwMutex.Lock()
	postList, ok := postsIndexMap[post.ParentId]
	//若字典中不存在该主题ID，则新建
	if !ok {
		postsIndexMap[post.ParentId] = []*Post{post}
	} else {
		postList = append(postList, post)
		postsIndexMap[post.ParentId] = postList
	}
	rwMutex.Unlock()
	return nil
}
