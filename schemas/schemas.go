package schemas

import (
  "gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role     string 
    Company  string
    location string
    remote   bool
    link     string
	salary   int64
}
