package services

import (
	"awesomeProject/internal/models"
)

func PredictLdl(data models.LdlRequest) ([]byte, error) {
	return SendRequest("https://apiml.labhub.online/api/v1/predict/ldl", data)
}
