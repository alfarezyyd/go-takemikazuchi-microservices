package dto

type CreateWorkerWalletDto struct {
	WalletType    string `form:"wallet_type" validate:"required,oneof=Bank DANA OVO GoPay LinkAja ShopeePay"`
	AccountName   string `form:"account_name" validate:"required,min=1,max=100"`
	AccountNumber string `form:"account_number" validate:"required,min=1,max=50"`
	BankName      string `form:"bank_name" validate:"required,min=1,max=100"`
}
