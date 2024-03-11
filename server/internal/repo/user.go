package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/nasdda/linkstore/internal/entity"
	"github.com/nasdda/linkstore/pkg/mongo"
)

type userRepo struct {
	db             mongo.Database
	userCollection mongo.Collection
}

func (ur *userRepo) create(ctx context.Context, user *entity.User) {
	user.ID = uuid.NewString()
	ur.userCollection.InsertOne(ctx, user)
}
