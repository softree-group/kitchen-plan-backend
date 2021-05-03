package infrastructure

import (
	"github.com/softree-group/kitchen-plan-backend/domain/repository"
	"github.com/softree-group/kitchen-plan-backend/infrastructure/persistance"
)

func NewRepositories() *repository.Repositories {
	return &repository.Repositories{
		Storage: persistance.NewPostgresStorage(),
	}
}
