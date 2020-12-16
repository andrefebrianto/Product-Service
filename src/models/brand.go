package models

import (
	"time"
)

//Brand datastructure for Brand domain
type Brand struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
