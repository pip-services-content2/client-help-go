package version1

type HelpArticleContentV1 struct {
	Language string            `json:"language"`
	Title    string            `json:"title"`
	Content  []*ContentBlockV1 `json:"content"`
}
