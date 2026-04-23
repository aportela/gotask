package models

type Attachment struct {
	ID        string `json:"id"`
	FileName  string `json:"fileName"`
	SizeBytes int64  `json:"sizeBytes"`
	CreatedBy User   `json:"createdBy"`
	CreatedAt int64  `json:"createdAt"`
}
