package worker

import (
	"go-takemikazuchi-api/internal/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
}

func NewRepository() *RepositoryImpl {
	return &RepositoryImpl{}
}

func (workerRepository RepositoryImpl) Store(gormTransaction *gorm.DB, workerModel *model.Worker) {
	gormTransaction.Create(workerModel)

}
