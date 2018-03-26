package projects

import (
	"../../core/structures"
	"gopkg.in/mgo.v2/bson"
)

type ProjectUser struct {
	ID bson.ObjectId     `json:"user,omitempty" bson:"_id,omitempty"`
	Permissions []int    `json:"permissions,omitempty" bson:"permissions,omitempty"`
	Creator bool         `json:"creator,omitempty" bson:"creator,omitempty"`
}

type Project struct {
	ID bson.ObjectId    `json:"id,omitempty" bson:"_id,omitempty"`
	Name string         `json:"name,omitempty" bson:"name,omitempty"`
	Description string  `json:"description,omitempty" bson:"description,omitempty"`
	Users []ProjectUser `json:"omitempty" bson:"users,omitempty"`
}

type ProjectCreation struct{
	Project Project      `json:"project,omitempty"`
	Session structures.Session `json:"session,omitempty"`
}