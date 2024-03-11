package usecase

import "github.com/nasdda/linkstore/internal/interfaces"

type link struct {
	linkRepo interfaces.LinkRepo
}

func NewLinkUseCase(linkRepo interfaces.LinkRepo) interfaces.Link {
	return &link{
		linkRepo: linkRepo,
	}
}
