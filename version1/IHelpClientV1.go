package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IHelpClientV1 interface {
	GetTopics(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpTopicV1], error)

	GetTopicById(ctx context.Context, correlationId string, topicId string) (*HelpTopicV1, error)

	CreateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error)

	UpdateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error)

	DeleteTopicById(ctx context.Context, correlationId string, topicId string) (*HelpTopicV1, error)

	GetArticles(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpArticleV1], error)

	GetRandomArticle(ctx context.Context, correlationId string, filter *data.FilterParams) (*HelpArticleV1, error)

	GetArticleById(ctx context.Context, correlationId string, articleId string) (*HelpArticleV1, error)

	CreateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error)

	UpdateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error)

	DeleteArticleById(ctx context.Context, correlationId string, articleId string) (*HelpArticleV1, error)
}
