package models

type Log struct {
	PaymentId  string   `json:"payment_id"`
	ErrorCodes []string `json:"error_codes"`
}

// LogWrapper log
type LogWrapper struct {
	Timestamp uint64 `json:"timestamp"`
	Log       Log    `json:"log"`
}
