package po

import (
	"time"

	"ddd/infrastructure/util/driver"
)

func init() {
	driver.DB.AutoMigrate(&Person{})
}

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
