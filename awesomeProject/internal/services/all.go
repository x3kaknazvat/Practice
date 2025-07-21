package services

import (
	"awesomeProject/internal/models"
	"fmt"
)

func PredictAll(age, gender int, chol, hdl, tg, ldl float64) models.AllPredictionResponse {
	hba1c := 4.3 + 0.03*float64(age) + 0.5*chol - 0.2*hdl + 0.1*tg
	ldlPred := 0.9*chol - 0.3*hdl
	hdlPred := 1.1 + 0.05*float64(age) - 0.2*tg + 0.1*ldl
	tgPred := 0.8 + 0.1*chol - 0.05*ldl + 0.03*hdl

	return models.AllPredictionResponse{
		Hba1c: fmt.Sprintf("%.1f mmol/mol", hba1c),
		Ldl:   fmt.Sprintf("%.1f mmol/L", ldlPred),
		Hdl:   fmt.Sprintf("%.1f mmol/L", hdlPred),
		Tg:    fmt.Sprintf("%.1f mmol/L", tgPred),
	}
}
