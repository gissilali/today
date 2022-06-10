package repositories

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Task       string
	IsDone     bool
	TaskListId *uint64
	AccountId  *uint32
}
