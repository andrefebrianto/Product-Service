package models

import (
	"time"
)

//Brand datastructure for Brand domain
type Brand struct {
	ID        string    `pg:"id,type:uuid,pk"`
	Name      string    `pg:"name,notnull"`
	CreatedAt time.Time `pg:"created_at,notnull,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,notnull,default:now()"`
	DeletedAt time.Time `pg:"deleted_at,soft_delete"`
}
