package models

import (
	"time"
)

type Resource struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Update      time.Time `json:"update"`
	Create      time.Time `json:"create"`
}
