package model

import "time"

type Post struct {
	Id      string    `json:"id,omitempty" bson:"_id"`
	Title   string    `json:"title,omitempty" bson:"title"`
	Content string    `json:"content,omitempty" bson:"content"`
	Created time.Time `json:"created,omitempty" bson:"created"`
	// updated
}
