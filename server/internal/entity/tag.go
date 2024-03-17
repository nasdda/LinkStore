package entity

type Tag struct {
	ID    OID    `json:"id" bson:"id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Color string `json:"color" bson:"color"`
}
