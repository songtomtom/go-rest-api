package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title    string `json:"title"`
	Archived bool   `json:"archived"`
	Tasks    []Task `json:"tasks"`
}

type Task struct {
	gorm.Model
	Title     string     `json:"title"`
	Priority  string     `json:"priority"`
	Deadline  *time.Time `json:"deadline"`
	Done      bool       `json:"done"`
	ProjectId uint       `json:"project_id"`
}
