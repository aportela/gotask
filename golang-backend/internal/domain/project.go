package domain

type Project struct {
	ID           string      `json:"id"`
	Key          string      `json:"key"`
	Summary      string      `json:"summary"`
	Description  *string     `json:"description"`
	CreatedBy    UserBase    `json:"createdBy"`
	CreatedAt    int64       `json:"createdAt"`
	UpdatedAt    *int64      `json:"updatedAt"`
	StartedAt    *int64      `json:"startedAt"`
	FinishedAt   *int64      `json:"finishedAt"`
	DueAt        *int64      `json:"dueAt"`
	Type         ProjectType `json:"type"`
	Participants []UserBase  `json:"participants"`
	//lead, asignee, priority, status
}
