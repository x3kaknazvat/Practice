package handler

import (
	"net/http"
	"strconv"

	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
)

func HandleTgPrediction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET allowed", http.StatusMethodNotAllowed)
		return
	}
	q := r.URL.Query()
	f := func(key string) float64 {
		v, _ := strconv.ParseFloat(q.Get(key), 64)
		return v
	}
	i := func(key string) int {
		v, _ := strconv.Atoi(q.Get(key))
		return v
	}

	req := models.TgRequest{
		Uid: q.Get("uid"), Age: i("age"), Gender: i("gender"),
		Rdw: f("rdw"), Wbc: f("wbc"), Rbc: f("rbc"),
		Hgb: f("hgb"), Hct: f("hct"), Mcv: f("mcv"),
		Mch: f("mch"), Mchc: f("mchc"), Plt: f("plt"),
		Neu: f("neu"), Eos: f("eos"), Bas: f("bas"),
		Lym: f("lym"), Mon: f("mon"), Soe: f("soe"),
		Chol: f("chol"), Glu: f("glu"),
	}

	resp, err := services.PredictTg(req)
	if err != nil {
		http.Error(w, "Request error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
