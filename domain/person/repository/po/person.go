package po

import (
	"time"
)

// Model 基本字段
type Model struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Person 人员模型
type Person struct {
	Model
	Name string
}
