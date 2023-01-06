package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-help-go/version1"
	"github.com/stretchr/testify/assert"
)

type HelpClientFixtureV1 struct {
	Client version1.IHelpClientV1

	HELP_TOPIC1 *version1.HelpTopicV1
	HELP_TOPIC2 *version1.HelpTopicV1

	HELP_ARTICLE1 *version1.HelpArticleV1
	HELP_ARTICLE2 *version1.HelpArticleV1
}

func NewHelpClientFixtureV1(client version1.IHelpClientV1) *HelpClientFixtureV1 {
	return &HelpClientFixtureV1{
		Client: client,
		HELP_TOPIC1: &version1.HelpTopicV1{
			Id:    "1",
			App:   "Test App 1",
			Title: map[string]string{"en": "Main topic"},
		},
		HELP_TOPIC2: &version1.HelpTopicV1{
			Id:       "2",
			ParentId: "1",
			App:      "Test App 1",
			Title:    map[string]string{"en": "Subtopic 1"},
			Popular:  true,
		},

		HELP_ARTICLE1: &version1.HelpArticleV1{
			Id:      "1",
			TopicId: "1",
			App:     "Test App 1",
			MinVer:  0,
			MaxVer:  9999,
			Status:  "new",
		},
		HELP_ARTICLE2: &version1.HelpArticleV1{
			Id:      "2",
			Tags:    []string{"TAG 1"},
			AllTags: []string{"tag1"},
			TopicId: "1",
			App:     "Test App 1",
			MinVer:  2,
			MaxVer:  9999,
			Status:  "new",
		},
	}
}

func (c *HelpClientFixtureV1) TestTopicsCrudOperations(t *testing.T) {
	// Create one topic
	topic1, err := c.Client.CreateTopic(context.Background(), "123", c.HELP_TOPIC1)
	assert.Nil(t, err)

	assert.NotNil(t, topic1)
	assert.Equal(t, topic1.Id, c.HELP_TOPIC1.Id)
	assert.Equal(t, topic1.App, c.HELP_TOPIC1.App)

	// Create another topic
	topic2, err := c.Client.CreateTopic(context.Background(), "123", c.HELP_TOPIC2)
	assert.Nil(t, err)

	assert.NotNil(t, topic2)
	assert.Equal(t, topic2.Id, c.HELP_TOPIC2.Id)
	assert.Equal(t, topic2.App, c.HELP_TOPIC2.App)

	// Get all topics
	page, err := c.Client.GetTopics(context.Background(), "123", nil, nil)
	assert.Nil(t, err)

	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)

	// Update the topic
	topic1.App = "New App 1"

	topic, err := c.Client.UpdateTopic(context.Background(), "123", topic1)
	assert.Nil(t, err)

	assert.NotNil(t, topic)
	assert.Equal(t, topic.App, "New App 1")
	assert.Equal(t, topic.Id, c.HELP_TOPIC1.Id)

	topic1 = topic

	// Delete topic
	_, err = c.Client.DeleteTopicById(context.Background(), "123", topic1.Id)
	assert.Nil(t, err)

	// Try to get delete topic
	topic, err = c.Client.GetTopicById(context.Background(), "123", topic1.Id)
	assert.Nil(t, err)
	assert.Nil(t, topic)

}

func (c *HelpClientFixtureV1) TestArticlesCrudOperations(t *testing.T) {
	// Create one article
	article1, err := c.Client.CreateArticle(context.Background(), "123", c.HELP_ARTICLE1)
	assert.Nil(t, err)

	assert.NotNil(t, article1)
	assert.Equal(t, article1.Id, c.HELP_ARTICLE1.Id)
	assert.Equal(t, article1.App, c.HELP_ARTICLE1.App)

	// Create another article
	article2, err := c.Client.CreateArticle(context.Background(), "123", c.HELP_ARTICLE2)
	assert.Nil(t, err)

	assert.NotNil(t, article2)
	assert.Equal(t, article2.Id, c.HELP_ARTICLE2.Id)
	assert.Equal(t, article2.App, c.HELP_ARTICLE2.App)

	// Get all articles
	page, err := c.Client.GetArticles(context.Background(), "123", nil, nil)
	assert.Nil(t, err)

	assert.NotNil(t, page)
	assert.Len(t, page.Data, 2)

	// Update the article
	article1.App = "New App 1"

	article, err := c.Client.UpdateArticle(context.Background(), "123", article1)
	assert.Nil(t, err)

	assert.NotNil(t, article)
	assert.Equal(t, article.App, "New App 1")
	assert.Equal(t, article.Id, c.HELP_TOPIC1.Id)

	article1 = article

	// Delete article
	_, err = c.Client.DeleteArticleById(context.Background(), "123", article1.Id)
	assert.Nil(t, err)

	// Try to get deleted article
	article, err = c.Client.GetArticleById(context.Background(), "123", article1.Id)
	assert.Nil(t, err)
	assert.Nil(t, article)

}
