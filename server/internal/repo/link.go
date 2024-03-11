package repo

import (
	"github.com/nasdda/linkstore/internal/interfaces"
	"github.com/nasdda/linkstore/pkg/mongo"
)

type linkRepo struct {
	db             mongo.Database
	userCollection mongo.Collection
}

func NewLinkRepo(db mongo.Database) interfaces.LinkRepo {
	userCollection := db.Collection("Link")
	return &linkRepo{
		db:             db,
		userCollection: userCollection,
	}
}
