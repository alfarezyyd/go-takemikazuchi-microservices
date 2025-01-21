package mapper

import (
	"go-takemikazuchi-api/internal/model"
)

func MapStringIntoWorkerResourceModel(workerId uint64, filesPath []string, typeFile []string, countFile int) []*model.WorkerResource {
	var workerResourcesModel []*model.WorkerResource
	for i := 0; i < countFile; i++ {
		workerResourceModel := model.WorkerResource{
			FilePath: filesPath[i],
			WorkerId: workerId,
			Type:     typeFile[i],
		}
		workerResourcesModel = append(workerResourcesModel, &workerResourceModel)
	}

	return workerResourcesModel
}
