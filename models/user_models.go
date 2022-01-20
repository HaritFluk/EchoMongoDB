package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id 			primitive.ObjectID 	`json:"id,omitempty"`
	Name 		string 				`json:"name,omitempty"`
	Location 	string 				`json:"location,omitempty"`
	Title 		string 				`json:"title,omitempty"`
}