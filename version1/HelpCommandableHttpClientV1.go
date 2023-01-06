package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type HelpCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewHelpCommandableHttpClientV1() *HelpCommandableHttpClientV1 {
	return &HelpCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/help"),
	}
}

func (c *HelpCommandableHttpClientV1) GetTopics(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpTopicV1], error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_topics", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*HelpTopicV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*HelpTopicV1]](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) GetTopicById(ctx context.Context, correlationId string, topicId string) (*HelpTopicV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"topic_id", topicId,
	)

	res, err := c.CallCommand(ctx, "get_topic_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpTopicV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) CreateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"topic", topic,
	)

	res, err := c.CallCommand(ctx, "create_topic", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpTopicV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) UpdateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"topic", topic,
	)

	res, err := c.CallCommand(ctx, "update_topic", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpTopicV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) DeleteTopicById(ctx context.Context, correlationId string, topicId string) (*HelpTopicV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"topic_id", topicId,
	)

	res, err := c.CallCommand(ctx, "delete_topic_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpTopicV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) GetArticles(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpArticleV1], error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	)

	res, err := c.CallCommand(ctx, "get_articles", correlationId, params)
	if err != nil {
		return *data.NewEmptyDataPage[*HelpArticleV1](), err
	}

	return clients.HandleHttpResponse[data.DataPage[*HelpArticleV1]](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) GetRandomArticle(ctx context.Context, correlationId string, filter *data.FilterParams) (*HelpArticleV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"filter", filter,
	)

	res, err := c.CallCommand(ctx, "get_random_article", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpArticleV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) GetArticleById(ctx context.Context, correlationId string, articleId string) (*HelpArticleV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"article_id", articleId,
	)

	res, err := c.CallCommand(ctx, "get_article_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpArticleV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) CreateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"article", article,
	)

	res, err := c.CallCommand(ctx, "create_article", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpArticleV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) UpdateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"article", article,
	)

	res, err := c.CallCommand(ctx, "update_article", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpArticleV1](res, correlationId)
}

func (c *HelpCommandableHttpClientV1) DeleteArticleById(ctx context.Context, correlationId string, articleId string) (*HelpArticleV1, error) {
	params := data.NewAnyValueMapFromTuples(
		"article_id", articleId,
	)

	res, err := c.CallCommand(ctx, "delete_article_by_id", correlationId, params)
	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*HelpArticleV1](res, correlationId)
}
