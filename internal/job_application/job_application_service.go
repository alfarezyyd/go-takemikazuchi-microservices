package job_application

import (
	"go-takemikazuchi-api/internal/job_application/dto"
)

type Service interface {
	HandleApply(applyJobApplicationDto *dto.ApplyJobApplicationDto)
}
