package service

import (
	"CommunityTopic_demo/repository"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}

func TestQueryPageInfo(t *testing.T) {
	pageData, _ := GetPageData(1)
	assert.NotEqual(t, nil, pageData)
	assert.Equal(t, 5, len(pageData.PostList))
}
