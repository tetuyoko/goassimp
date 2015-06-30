package models

import "time"

type Template struct {
    ID           uint `gorm:"primary_key"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
}
