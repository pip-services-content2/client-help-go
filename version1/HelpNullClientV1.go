package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type HelpNullClientV1 struct {
}

func NewHelpNullClientV1() *HelpNullClientV1 {
	return &HelpNullClientV1{}
}

func (c *HelpNullClientV1) GetTopics(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpTopicV1], error) {
	return *data.NewEmptyDataPage[*HelpTopicV1](), nil
}

func (c *HelpNullClientV1) GetTopicById(ctx context.Context, correlationId string, topicId string) (*HelpTopicV1, error) {
	return nil, nil
}

func (c *HelpNullClientV1) CreateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error) {
	return topic, nil
}

func (c *HelpNullClientV1) UpdateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error) {
	return topic, nil
}

func (c *HelpNullClientV1) DeleteTopicById(ctx context.Context, correlationId string, topicId string) (*HelpTopicV1, error) {
	return nil, nil
}

func (c *HelpNullClientV1) GetArticles(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpArticleV1], error) {
	return *data.NewEmptyDataPage[*HelpArticleV1](), nil
}

func (c *HelpNullClientV1) GetRandomArticle(ctx context.Context, correlationId string, filter *data.FilterParams) (*HelpArticleV1, error) {
	return nil, nil
}

func (c *HelpNullClientV1) GetArticleById(ctx context.Context, correlationId string, articleId string) (*HelpArticleV1, error) {
	return nil, nil
}

func (c *HelpNullClientV1) CreateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error) {
	return article, nil
}

func (c *HelpNullClientV1) UpdateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error) {
	return article, nil
}

func (c *HelpNullClientV1) DeleteArticleById(ctx context.Context, correlationId string, articleId string) (*HelpArticleV1, error) {
	return nil, nil
}
