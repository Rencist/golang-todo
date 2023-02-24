package entity

import "time"

type Todo struct {
	ID        	uint64 		`json:"id"`
	Todo 		string 		`json:"todo"`
	IsDone 		bool 		`json:"is_done" default:"false"`
	CreatedAt 	time.Time 	`json:"created_at" default:"current_timestamp"`
}