package dto

type CreateTransactionDto struct {
	JobId       uint64 `json:"job_id" validate:"required,gte=1"`
	ApplicantId uint64 `json:"applicant_id" validate:"required,gte=1"`
}

type TransactionNotificationDto struct {
	TransactionId          string  `json:"transaction_id"`
	TransactionTime        string  `json:"transaction_time"`
	TransactionStatus      string  `json:"transaction_status"`
	StatusMessage          string  `json:"status_message"`
	StatusCode             string  `json:"status_code"`
	SignatureKey           string  `json:"signature_key"`
	PaymentType            string  `json:"payment_type"`
	OrderId                string  `json:"order_id"`
	MerchantId             string  `json:"merchant_id"`
	MaskedCard             string  `json:"masked_card"`
	GrossAmount            string  `json:"gross_amount"`
	FraudStatus            string  `json:"fraud_status"`
	Eci                    string  `json:"eci"`
	Currency               string  `json:"currency"`
	ChannelResponseMessage string  `json:"channel_response_message"`
	ChannelResponseCode    *string `json:"channel_response_code"`
	CardType               string  `json:"card_type"`
	Bank                   string  `json:"bank"`
	ApprovalCode           string  `json:"approval_code"`
}
