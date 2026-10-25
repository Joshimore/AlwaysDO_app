package todo

import (

	"time"
)

type Task struct {
	ID          int
	Category    string
	Title       string
	Done        bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}
