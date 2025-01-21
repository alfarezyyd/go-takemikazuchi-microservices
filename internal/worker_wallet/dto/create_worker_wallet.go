package dto

type CreateWorkerWalletDto struct {
	WalletType    string `json:"wallet_type" validate:"required,oneof=Bank DANA OVO GoPay LinkAja ShopeePay"`
	AccountName   string `json:"account_name" validate:"required,min=1,max=100"`
	AccountNumber string `json:"account_number" validate:"required,min=1,max=50"`
	BankName      string `json:"bank_name" validate:"required,min=1,max=100"`
}
