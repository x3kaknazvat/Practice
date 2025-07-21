package models

type AllPredictionResponse struct {
	Hba1c string `json:"hba1c"`
	Ldl   string `json:"ldl"`
	Hdl   string `json:"hdl"`
	Tg    string `json:"tg"`
}
