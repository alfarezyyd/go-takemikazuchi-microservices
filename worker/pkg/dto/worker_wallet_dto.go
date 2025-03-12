package dto

type CreateWorkerWalletDto struct {
	WalletType    string `form:"wallet_type" validate:"required,oneof=Bank DANA OVO GoPay LinkAja ShopeePay" mapstructure:"WalletType"`
	AccountName   string `form:"account_name" validate:"required,min=1,max=100" mapstructure:"AccountName"`
	AccountNumber string `form:"account_number" validate:"required,min=1,max=50"`
	BankName      string `form:"bank_name" validate:"required,min=1,max=100" mapstructure:"BankName"`
}

type ResponseWorkerWalletDto struct {
	ID            uint64
	WorkerID      uint64
	WalletType    string
	AccountName   string
	AccountNumber string
	BankName      string
	IsPrimary     bool
	CreatedAt     string `mapstructure:"-"`
	UpdatedAt     string `mapstructure:"-"`
}
