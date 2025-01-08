package job_application

import "go-takemikazuchi-api/job_application/dto"

type Service interface {
	HandleApply(applyJobApplicationDto *dto.ApplyJobApplicationDto)
}
