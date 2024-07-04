package models

type SendOTPResponse struct {
	ResponseCode int     `json:"responseCode"`
	Message      string  `json:"message"`
	Data         OTPData `json:"data"`
}

type OTPData struct {
	VerificationID string `json:"verificationId"`
	MobileNumber   string `json:"mobileNumber"`
	ResponseCode   string `json:"responseCode"`
	ErrorMessage   string `json:"errorMessage"`
	Timeout        string `json:"timeout"`
	SmCLI          string `json:"smCLI"`
	TransactionID  string `json:"transactionId"`
}

type ValidateOTPResponse struct {
	ResponseCode int             `json:"responseCode"`
	Message      string          `json:"message"`
	Data         ValidateOTPData `json:"data"`
}

type ValidateOTPData struct {
	VerificationId     int    `json:"verificationId"`
	MobileNumber       string `json:"mobileNumber"`
	ResponseCode       string `json:"responseCode"`
	ErrorMessage       string `json:"errorMessage"`
	VerificationStatus string `json:"verificationStatus"`
	AuthToken          string `json:"authToken"`
	TransactionID      string `json:"transactionId"`
}
