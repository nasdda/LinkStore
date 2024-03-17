package entity

const (
	PUBLIC  = "public"
	PRIVATE = "private"
)

type Collection struct {
	ID    OID      `json:"id" bson:"id,omitempty"`
	Name  string   `json:"name" bson:"name"`
	Links []string `json:"links" bson:"links"`
}
