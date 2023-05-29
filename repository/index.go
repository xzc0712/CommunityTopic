package repository

import (
	"bufio"
	"encoding/json"
	"os"
)

var topicIndexMap map[int]*Topic
var postsIndexMap map[int][]*Post

func Init(filepath string) error {
	if err := initTopicDoc(filepath); err != nil {
		return err
	}
	if err := initPostDoc(filepath); err != nil {
		return err
	}
	return nil
}

func initPostDoc(filepath string) error {
	open, err := os.Open(filepath + "post")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	postsTmpPost := make(map[int][]*Post)
	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		err := json.Unmarshal([]byte(text), &post)
		if err != nil {
			return err
		}
		posts, ok := postsTmpPost[post.ParentId]
		if !ok {
			postsTmpPost[post.ParentId] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postsTmpPost[post.ParentId] = posts
	}
	postsIndexMap = postsTmpPost
	return nil
}

func initTopicDoc(filepath string) error {
	open, err := os.Open(filepath + "topic")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(open)
	topicTmpMap := make(map[int]*Topic)
	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		err := json.Unmarshal([]byte(text), &topic)
		if err != nil {
			return err
		}
		topicTmpMap[topic.Id] = &topic
	}
	topicIndexMap = topicTmpMap
	return nil
}
