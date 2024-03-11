package entity

type Link struct {
	ID          string `json:"id" bson:"id,omitempty"`
	Name        string `json:"name" bson:"name"`
	URL         string `json:"url" bson:"url"`
	Description string `json:"descriptions" bson:"descriptions"`
}
