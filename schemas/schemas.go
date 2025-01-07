package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role    string
	Company string
	Remote  bool
	Link    string
	Salary  int64
}

type OpeningResponse struct {
	ID       uint      `json: "id"`
	CreateAt time.Time `json: "create_at"`
	UpdateAt time.Time `json: "update_at"`
	DeleteAt time.Time `json: "delete_at, omitempty"`
	Role     string    `json: "role"`
	Company  string    `json: "company"`
	Remote   bool      `json: "remote"`
	Link     string    `json: "link"`
	Salary   int64     `json: "salary"`
}
