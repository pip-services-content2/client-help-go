package version1

import (
	"context"

	aclients "github.com/pip-services-content2/client-attachments-go/version1"
)

type AttachmentsConnector struct {
	attachmentsClient aclients.IAttachmentsClientV1
}

func NewAttachmentsConnector(client aclients.IAttachmentsClientV1) *AttachmentsConnector {
	return &AttachmentsConnector{
		attachmentsClient: client,
	}
}

func (c *AttachmentsConnector) extractAttachmentIds(article *HelpArticleV1) []string {
	ids := make([]string, 0)

	if article == nil {
		return ids
	}

	if article.Content == nil {
		article.Content = make([]*HelpArticleContentV1, 0)
	}

	for _, content := range article.Content {
		if content.Content == nil {
			content.Content = make([]*ContentBlockV1, 0)
		}
		for _, block := range content.Content {
			if block.PicIds == nil {
				block.PicIds = make([]string, 0)
			}
			if block.Docs == nil {
				block.Docs = make([]*DocumentV1, 0)
			}

			ids = append(ids, block.PicIds...)
			for _, doc := range block.Docs {
				ids = append(ids, doc.FileId)
			}
		}
	}

	return ids
}

func (c *AttachmentsConnector) AddAttachments(ctx context.Context, correlationId string, article *HelpArticleV1) error {
	if c.attachmentsClient == nil || article == nil {
		return nil
	}

	ids := c.extractAttachmentIds(article)
	reference := aclients.NewReferenceV1(article.Id, "help-article", "")
	_, err := c.attachmentsClient.AddAttachments(ctx, correlationId, reference, ids)
	return err
}

func (c *AttachmentsConnector) UpdateAttachments(ctx context.Context, correlationId string, oldArticle *HelpArticleV1, newArticle *HelpArticleV1) error {
	if c.attachmentsClient == nil || oldArticle == nil || newArticle == nil {
		return nil
	}

	oldIds := c.extractAttachmentIds(oldArticle)
	newIds := c.extractAttachmentIds(newArticle)
	reference := aclients.NewReferenceV1(newArticle.Id, "help-article", "")
	_, err := c.attachmentsClient.UpdateAttachments(ctx, correlationId, reference, oldIds, newIds)
	return err
}

func (c *AttachmentsConnector) RemoveAttachments(ctx context.Context, correlationId string, article *HelpArticleV1) error {
	if c.attachmentsClient == nil || article == nil {
		return nil
	}

	ids := c.extractAttachmentIds(article)
	reference := aclients.NewReferenceV1(article.Id, "help-article", "")
	_, err := c.attachmentsClient.RemoveAttachments(ctx, correlationId, reference, ids)
	return err
}
