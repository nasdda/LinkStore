package entity

type User struct {
	ID          string   `json:"id" bson:"id,omitempty"`
	FirstName   string   `json:"firstName" bson:"firstName"`
	LastName    string   `json:"lastName" bson:"firstName"`
	Email       string   `json:"email" bson:"firstName"`
	Collections []string `json:"collections" bson:"collections"`
}
