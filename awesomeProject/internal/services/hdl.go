package services

import (
	"awesomeProject/internal/models"
)

func PredictHdl(data models.HdlRequest) ([]byte, error) {
	return SendRequest("https://apiml.labhub.online/api/v1/predict/hdl", data)
}
