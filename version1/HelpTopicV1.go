package version1

type HelpTopicV1 struct {
	Id       string            `json:"id"`
	ParentId string            `json:"parent_id"`
	App      string            `json:"app"`
	Index    int               `json:"index"`
	Title    map[string]string `json:"title"`
	Popular  bool              `json:"popular"`

	// Custom fields
	CustomHdr any `json:"custom_hdr"`
	CustomDat any `json:"custom_dat"`
}
