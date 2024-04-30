package repository

import "context"

type Repository struct {
	ctx context.Context
}

func Create(ctx context.Context) *Repository {
	return &Repository{ctx: ctx}
}

