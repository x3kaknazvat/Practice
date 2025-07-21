package handler

import (
	"net/http"
	"strconv"

	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
)

func HandleHba1cPrediction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query()
	parseF := func(key string) float64 {
		v, _ := strconv.ParseFloat(q.Get(key), 64)
		return v
	}
	parseI := func(key string) int {
		v, _ := strconv.Atoi(q.Get(key))
		return v
	}

	req := models.Hba1cRequest{
		Uid:    q.Get("uid"),
		Age:    parseI("age"),
		Gender: parseI("gender"),
		Rdw:    parseF("rdw"), Wbc: parseF("wbc"), Rbc: parseF("rbc"),
		Hgb: parseF("hgb"), Hct: parseF("hct"), Mcv: parseF("mcv"),
		Mch: parseF("mch"), Mchc: parseF("mchc"), Plt: parseF("plt"),
		Neu: parseF("neu"), Eos: parseF("eos"), Bas: parseF("bas"),
		Lym: parseF("lym"), Mon: parseF("mon"), Soe: parseF("soe"),
		Chol: parseF("chol"), Glu: parseF("glu"),
	}

	resp, err := services.PredictHba1c(req)
	if err != nil {
		http.Error(w, "Request error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
