package services

import (
	"awesomeProject/internal/models"
)

func PredictTg(data models.TgRequest) ([]byte, error) {
	return SendRequest("https://apiml.labhub.online/api/v1/predict/tg", data)
}
