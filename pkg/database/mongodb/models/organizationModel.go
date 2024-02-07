// Organization model
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Organization model with name and description and an array of users containing the user name and email and access level
type Organization struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Users       []UserDetails      `bson:"users"`
}

// declare the data in users array
type UserDetails struct {
	Name         string `bson:"name"`
	Email        string `bson:"email"`
	Access_level string `bson:"access"`
}
