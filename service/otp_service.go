package service

import (
	"bytes"
	"encoding/json"
	"message-central-v3/auth"
	"message-central-v3/config"
	"message-central-v3/models"
	"message-central-v3/utils"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2/log"
)

func SendOTP(request models.SendOTPRequest) (*models.SendOTPResponse, error) {
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		log.Error("Error parsing base URL:", err)
		return nil, err
	}

	// Construct the full URL with query parameters
	params := url.Values{}
	params.Add("countryCode", request.CountryCode)
	params.Add("flowType", request.FlowType)
	params.Add("mobileNumber", request.MobileNumber)

	baseURL.Path += "/verification/v3/send"
	baseURL.RawQuery = params.Encode()

	authToken := auth.GetAuthToken()
	log.Info("authToken")
	log.Info(authToken)
	if authToken == "" {
		log.Info("Unauthorized: Missing authToken")
	}

	log.Info("Sending OTP to URL:", baseURL.String())

	// Make the HTTP request
	req, err := http.NewRequest("POST", baseURL.String(), bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}
	req.Header.Set("authToken", authToken)

	res, err := utils.HttpClient.Do(req)
	if err != nil {
		log.Error("HTTP request error:", err)
		return nil, err
	}
	defer res.Body.Close()

	var response models.SendOTPResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error("Error decoding response:", err)
		return nil, err
	}

	return &response, nil
}

func ValidateOTP(request models.ValidateOTPRequest) (*models.ValidateOTPResponse, error) {
	url := config.BaseURL + "/verification/v3/validateOtp?verificationId=" + request.VerificationId + "&code=" + request.Code

	authToken := auth.GetAuthToken()
	if authToken == "" {
		log.Error("Unauthorized: Missing authToken")
	}
	log.Info(authToken)

	// Make the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	req.Header.Set("authToken", authToken)

	res, err := utils.HttpClient.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()

	var response models.ValidateOTPResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Error(err)
		return nil, err
	}

	return &response, nil
}
