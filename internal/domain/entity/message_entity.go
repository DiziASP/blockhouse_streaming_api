package entity

import "time"

type MessageEntity struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
