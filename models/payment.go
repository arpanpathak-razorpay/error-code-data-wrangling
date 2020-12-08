package models

type Payment struct {
	GatewayResponseStatus   string `json:"gateway_response_status"`
	Status                  string `json:"mozart_status"`
	ErrorDescription        string `json:"error_description"`
	GatewayErrorCode        string `json:"gateway_error_code"`
	GatewayErrorDescription string `json:"gateway_error_description"`
	GatewayStatusCode       string `json:"gateway_status_code"`
	InternalErrorCode       string `json:"internal_error_code"`
	Environment             string `json:"environment"`
	Mode                    string `json:"mode"`
	PaymentID               string `json:"payment_id"`
	TimeStamp               string `json:"timestamp"`
}
