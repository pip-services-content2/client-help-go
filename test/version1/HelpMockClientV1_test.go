package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-help-go/version1"
)

type HelpMockClientV1 struct {
	client  *version1.HelpMockClientV1
	fixture *HelpClientFixtureV1
}

func newHelpMockClientV1() *HelpMockClientV1 {
	return &HelpMockClientV1{}
}

func (c *HelpMockClientV1) setup(t *testing.T) *HelpClientFixtureV1 {
	c.client = version1.NewHelpMockClientV1()
	c.fixture = NewHelpClientFixtureV1(c.client)
	return c.fixture
}

func (c *HelpMockClientV1) teardown(t *testing.T) {
	topics, _ := c.client.GetTopics(context.Background(), "123", nil, nil)
	for _, topic := range topics.Data {
		c.client.DeleteTopicById(context.Background(), "123", topic.Id)
	}

	articles, _ := c.client.GetArticles(context.Background(), "123", nil, nil)
	for _, article := range articles.Data {
		c.client.DeleteArticleById(context.Background(), "123", article.Id)
	}
}

func TestMockCrudOperations(t *testing.T) {
	c := newHelpMockClientV1()

	fixture := c.setup(t)
	fixture.TestTopicsCrudOperations(t)
	c.teardown(t)

	fixture = c.setup(t)
	fixture.TestArticlesCrudOperations(t)
	c.teardown(t)
}
