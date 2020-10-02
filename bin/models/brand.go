package models

import (
	"time"
)

//Brand datastructure for Brand domain
type Brand struct {
	id        string
	name      string
	createdAt time.Time
	updatedAt time.Time
}
