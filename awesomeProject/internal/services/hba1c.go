package services

import (
	"awesomeProject/internal/models"
)

func PredictHba1c(data models.Hba1cRequest) ([]byte, error) {
	return SendRequest("https://apiml.labhub.online/api/v1/predict/hba1c", data)
}
