package model

type Category struct {
	ID          string `json:"id"          bson:"_id,omitempty"`
	Name        string `json:"name"        bson:"name,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
}
