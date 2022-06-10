package repositories

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	task       string
	isDone     bool
	taskListId *uint64
	accountId  *uint32
}
