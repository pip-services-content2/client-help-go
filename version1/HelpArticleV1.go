package version1

import "time"

type HelpArticleV1 struct {
	// Identification
	Id      string `json:"id"`
	TopicId string `json:"topic_id"`
	App     string `json:"app"`
	Index   int    `json:"index"`
	MinVer  int    `json:"min_ver"`
	MaxVer  int    `json:"max_ver"`

	// Auto-generated fields
	CreateTime time.Time `json:"create_time"`

	// Content
	Content []*HelpArticleContentV1 `json:"content"`

	// Search
	Tags    []string `json:"tags"`
	AllTags []string `json:"all_tags"`

	// Status
	Status string `json:"status"`

	// Custom fields
	CustomHdr any `json:"custom_hdr"`
	CustomDat any `json:"custom_dat"`
}
