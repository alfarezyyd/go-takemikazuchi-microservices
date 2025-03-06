package dto

import "go-takemikazuchi-microservices/internal/model"

func ConstructTransactionModel(jobApplicationModel *model.JobApplication, jobModel *model.Job, transactionModel *model.Transaction) {
	transactionModel.JobID = jobModel.ID
	transactionModel.PayerID = jobModel.UserId
	transactionModel.PayeeID = jobApplicationModel.ApplicantId
	transactionModel.Amount = jobModel.Price
}
