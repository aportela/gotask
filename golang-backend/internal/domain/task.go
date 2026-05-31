package domain

type TaskBase struct {
	ID   string
	Slug string
}

type Task struct {
	ID          string
	Slug        string
	Summary     string
	Description string
	Status      TaskStatus
	Priority    TaskPriority
	CreatedBy   UserBase
	CreatedAt   int64
	StartedAt   *int64
	FinishedAt  *int64
	DueAt       *int64
	Tags        []string
	Notes       []Note
	Attachments []Attachment
	LinkedTasks []TaskBase
}
