// api-seguridad/resources/area_chiefs/application/delete_usecase.go
package application

import (
	"context"
	"errors"
	"api-seguridad/resources/area_chiefs/domain/repository"
)

type DeleteAreaChiefUseCase struct {
	repo repository.AreaChiefRepository
}

func NewDeleteAreaChiefUseCase(repo repository.AreaChiefRepository) *DeleteAreaChiefUseCase {
	return &DeleteAreaChiefUseCase{repo: repo}
}

func (uc *DeleteAreaChiefUseCase) Execute(ctx context.Context, id uint) error {
	if id == 0 {
		return errors.New("invalid chief ID")
	}

	existing, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if existing == nil || existing.IsDeleted() {
		return errors.New("area chief not found")
	}

	return uc.repo.Delete(ctx, id)
}