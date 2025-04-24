//api-seguridad/resources/police/application/search_by_name_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/police/domain/entities"
	"api-seguridad/resources/police/domain/repository"
)

type SearchPoliceByNameUseCase struct {
	repo repository.PoliceRepository
}

func NewSearchPoliceByNameUseCase(repo repository.PoliceRepository) *SearchPoliceByNameUseCase {
	return &SearchPoliceByNameUseCase{repo: repo}
}

func (uc *SearchPoliceByNameUseCase) Execute(ctx context.Context, name string) ([]*entities.Police, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	return uc.repo.SearchByName(ctx, name)
}