package dto

type CreateWithdrawalDto struct {
	WalletId uint64 `json:"wallet_id"`
	Amount   int64  `json:"amount"`
}
