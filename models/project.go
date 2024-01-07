package models

import "time"

type Project struct {
	ID           string    `json:"id" bson:"_id"`
	Title        string    `json:"title" bson:"title"`
	Author       string    `json:"author" bson:"author"`
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
	Tasks        []Task    `json:"tasks" bson:"tasks"`
}
