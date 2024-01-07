package models

import "time"

type Task struct {
	ID           string    `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string    `json:"title" bson:"title"`
	Author       string    `json:"author" bson:"author"`
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	Description  string    `json:"description" bson:"description"`
}
