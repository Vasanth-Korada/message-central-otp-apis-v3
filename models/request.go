package models

type SendOTPRequest struct {
	CountryCode  string `json:"countryCode"`
	MobileNumber string `json:"mobileNumber"`
	FlowType     string `json:"flowType"`
	OTPLength    int    `json:"otpLength"`
}

type ValidateOTPRequest struct {
	VerificationId string `query:"verificationId"`
	Code           string `query:"code"`
	LangID         string `query:"langId"`
}
