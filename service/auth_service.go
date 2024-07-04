package service

import (
	"encoding/json"
	"message-central-v3/config"
	"message-central-v3/utils"
)

func GenerateToken() (string, error) {
	url := config.BaseURL + "/auth/v1/authentication/token?customerId=" + config.CustomerID + "&key=" + config.Key + "&scope=NEW"

	res, err := utils.HttpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var response struct {
		Status int    `json:"status"`
		Token  string `json:"token"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return "", err
	}

	return response.Token, nil
}
