package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title    string     `json:"title"`
	Priority string     `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline *time.Time `gorm:"default:null" json:"deadline"`
	Done     bool       `json:"done"`
	// ProjectID uint       `json:"project_id"`
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) Undo() {
	t.Done = false
}
