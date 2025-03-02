package dto

import jobDto "go-takemikazuchi-microservices/internal/job/dto"

type CategoryResponseDto struct {
	ID          uint64                  `json:"id,omitempty"`
	Name        string                  `json:"name,omitempty"`
	Description string                  `json:"description,omitempty"`
	Jobs        []jobDto.JobResponseDto `json:"jobs,omitempty"`
}
