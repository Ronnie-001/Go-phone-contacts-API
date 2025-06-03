package models

import (
	"time"
)

type Contact struct {
	name string
	number uint64
	dateAdded time.Time
	isFavorite bool
	notes string
}


