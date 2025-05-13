package helper

import (
	"github.com/zercle/kku-qf-services/pkg/domain"
	"gorm.io/gorm"
)

type Transaction = func(tx *gorm.DB) error

func ExecuteTransaction(
	repo domain.CommonRepository,
	actions ...Transaction,
) error {

	tx, err := repo.Begin()
	if err != nil {
		return err
	}

	for i := range actions {
		err := actions[i](tx)
		if err != nil {
			repo.Rollback(tx)
			return err
		}
	}

	if err := repo.Commit(tx); err != nil {
		repo.Rollback(tx)
		return err
	}

	return nil
}
