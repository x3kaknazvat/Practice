package services

import (
	"awesomeProject/internal/models"
)

func PredictFerr(data models.FerrRequest) ([]byte, error) {
	return SendRequest("https://apiml.labhub.online/api/v1/predict/ferr", data)
}
