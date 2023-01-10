package version1

import (
	"context"
	"time"

	aclients "github.com/pip-services-content2/client-attachments-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
	"github.com/pip-services3-gox/pip-services3-commons-gox/random"
)

type HelpMockClientV1 struct {
	attachmentsConnector *AttachmentsConnector
	attachmentsClient    *aclients.AttachmentsMockClientV1

	topics   []*HelpTopicV1
	articles []*HelpArticleV1
}

func NewHelpMockClientV1() *HelpMockClientV1 {
	c := &HelpMockClientV1{}
	c.attachmentsClient = aclients.NewAttachmentsMockClientV1()
	c.attachmentsConnector = NewAttachmentsConnector(c.attachmentsClient)
	c.topics = make([]*HelpTopicV1, 0)
	c.articles = make([]*HelpArticleV1, 0)

	return c
}

func (c *HelpMockClientV1) contains(array1 []string, array2 []string) bool {
	if array1 == nil || array2 == nil {
		return false
	}

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			if i < len(array2) {
				if array1[i] == array2[i] {
					return true
				}
			} else {
				break
			}
		}
	}

	return false
}

func (c *HelpMockClientV1) composeTopicsFilter(filter *data.FilterParams) func(item *HelpTopicV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	id, idOk := filter.GetAsNullableString("id")
	parentId, parentIdOk := filter.GetAsNullableString("parent_id")
	app, appOk := filter.GetAsNullableString("app")
	popular, popularOk := filter.GetAsNullableBoolean("popular")

	return func(item *HelpTopicV1) bool {
		if idOk && id != item.Id {
			return false
		}
		if parentIdOk && parentId != item.ParentId {
			return false
		}
		if appOk && app != item.App {
			return false
		}
		if popularOk && popular != item.Popular {
			return false
		}
		return true
	}
}

func (c *HelpMockClientV1) composeArticlesFilter(filter *data.FilterParams) func(item *HelpArticleV1) bool {
	if filter == nil {
		filter = data.NewEmptyFilterParams()
	}

	id, idOk := filter.GetAsNullableString("id")
	topicId, topicIdOk := filter.GetAsNullableString("topic_id")
	app, appOk := filter.GetAsNullableString("app")
	version, versionOk := filter.GetAsNullableInteger("version")
	status, statusOk := filter.GetAsNullableString("status")
	tagsString := filter.GetAsString("tags")
	tags := make([]string, 0)

	if tagsString != "" {
		tags = data.TagsProcessor.CompressTags([]string{tagsString})
	}

	return func(item *HelpArticleV1) bool {
		if idOk && id != item.Id {
			return false
		}
		if topicIdOk && topicId != item.TopicId {
			return false
		}
		if appOk && app != item.App {
			return false
		}
		if versionOk && (version < item.MaxVer || version > item.MaxVer) {
			return false
		}
		if statusOk && status != item.Status {
			return false
		}
		if len(tags) > 0 && !c.contains(item.AllTags, tags) {
			return false
		}
		return true
	}
}

func (c *HelpMockClientV1) GetTopics(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpTopicV1], error) {
	filterFunc := c.composeTopicsFilter(filter)

	items := make([]*HelpTopicV1, 0)
	for _, v := range c.topics {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.topics)), nil
}

func (c *HelpMockClientV1) GetTopicById(ctx context.Context, correlationId string, topicId string) (result *HelpTopicV1, err error) {
	for _, v := range c.topics {
		if v.Id == topicId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, err
}

func (c *HelpMockClientV1) CreateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error) {
	buf := *topic
	c.topics = append(c.topics, &buf)
	return topic, nil
}

func (c *HelpMockClientV1) UpdateTopic(ctx context.Context, correlationId string, topic *HelpTopicV1) (*HelpTopicV1, error) {
	if topic == nil {
		return nil, nil
	}

	var index = -1
	for i, v := range c.topics {
		if v.Id == topic.Id {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}

	buf := *topic
	c.topics[index] = &buf
	return topic, nil
}

func (c *HelpMockClientV1) DeleteTopicById(ctx context.Context, correlationId string, topicId string) (*HelpTopicV1, error) {
	var index = -1
	for i, v := range c.topics {
		if v.Id == topicId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	var item = c.topics[index]
	if index < len(c.topics) {
		c.topics = append(c.topics[:index], c.topics[index+1:]...)
	} else {
		c.topics = c.topics[:index]
	}
	return item, nil
}

func (c *HelpMockClientV1) GetArticles(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (data.DataPage[*HelpArticleV1], error) {
	filterFunc := c.composeArticlesFilter(filter)

	items := make([]*HelpArticleV1, 0)
	for _, v := range c.articles {
		item := *v
		if filterFunc(&item) {
			items = append(items, &item)
		}
	}
	return *data.NewDataPage(items, len(c.articles)), nil
}

func (c *HelpMockClientV1) GetRandomArticle(ctx context.Context, correlationId string, filter *data.FilterParams) (*HelpArticleV1, error) {
	filterFunc := c.composeArticlesFilter(filter)

	items := make([]*HelpArticleV1, 0)
	for _, v := range c.articles {
		item := v
		if filterFunc(item) {
			items = append(items, item)
		}
	}

	buf := *items[random.Integer.Next(0, len(items))]
	return &buf, nil
}

func (c *HelpMockClientV1) GetArticleById(ctx context.Context, correlationId string, articleId string) (result *HelpArticleV1, err error) {
	for _, v := range c.articles {
		if v.Id == articleId {
			buf := *v
			result = &buf
			break
		}
	}
	return result, err
}

func (c *HelpMockClientV1) CreateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error) {
	article.CreateTime = time.Now()
	article.AllTags = data.TagsProcessor.ExtractHashTags(
		"#content",
	)

	buf := *article
	c.articles = append(c.articles, &buf)

	err := c.attachmentsConnector.AddAttachments(ctx, correlationId, &buf)

	return article, err
}

func (c *HelpMockClientV1) UpdateArticle(ctx context.Context, correlationId string, article *HelpArticleV1) (*HelpArticleV1, error) {
	if article == nil {
		return nil, nil
	}
	var oldArticle *HelpArticleV1

	article.AllTags = data.TagsProcessor.ExtractHashTags(
		"#content",
	)

	for _, v := range c.articles {
		if v.Id == article.Id {
			oldArticle = v
			break
		}
	}

	if oldArticle == nil {
		return nil, errors.NewNotFoundError(
			correlationId,
			"ARTICLE_NOT_FOUND",
			"Help article "+article.Id+" was not found",
		).WithDetails("article_id", article.Id)
	}

	index := -1
	for i, v := range c.articles {
		if v.Id == article.Id {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, nil
	}

	buf := *article
	c.articles[index] = &buf

	err := c.attachmentsConnector.UpdateAttachments(ctx, correlationId, oldArticle, &buf)

	return article, err
}

func (c *HelpMockClientV1) DeleteArticleById(ctx context.Context, correlationId string, articleId string) (*HelpArticleV1, error) {
	var index = -1
	for i, v := range c.articles {
		if v.Id == articleId {
			index = i
			break
		}
	}

	if index < 0 {
		return nil, nil
	}
	var oldArticle = c.articles[index]
	if index < len(c.articles) {
		c.articles = append(c.articles[:index], c.articles[index+1:]...)
	} else {
		c.articles = c.articles[:index]
	}

	err := c.attachmentsConnector.RemoveAttachments(ctx, correlationId, oldArticle)
	return oldArticle, err
}
