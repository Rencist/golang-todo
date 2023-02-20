package entity

import "time"

type Todo struct {
	ID        	uint64 		`json:"id"`
	Todo 		string 		`json:"todo"`
	CreatedAt 	time.Time 	`json:"created_at" default:"current_timestamp"`
}