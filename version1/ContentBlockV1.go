package version1

import "time"

type ContentBlockV1 struct {
	Type      string             `json:"type"`
	Text      string             `json:"text"`
	Checklist []*ChecklistItemV1 `json:"check_list"`
	LocName   string             `json:"loc_name"`
	LocPos    map[string]any     `json:"loc_pos"` // GeoJson
	Start     time.Time          `json:"start"`
	End       time.Time          `json:"end"`
	AllDay    bool               `json:"all_day"`
	PicIds    []string           `json:"pic_ids"`
	Docs      []*DocumentV1      `json:"docs"`
	EmbedType string             `json:"embed_type"`
	EmbedUri  string             `json:"embed_uri"`
	Custom    any                `json:"custom"`
}
