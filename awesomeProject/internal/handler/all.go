package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"awesomeProject/internal/models"
	"awesomeProject/internal/services"
)

func HandleAllPrediction(w http.ResponseWriter, r *http.Request) {
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

	// HBA1c
	hreq := models.Hba1cRequest{Uid: q.Get("uid"), Age: i("age"), Gender: i("gender"),
		Rdw: f("rdw"), Wbc: f("wbc"), Rbc: f("rbc"),
		Hgb: f("hgb"), Hct: f("hct"), Mcv: f("mcv"),
		Mch: f("mch"), Mchc: f("mchc"), Plt: f("plt"),
		Neu: f("neu"), Eos: f("eos"), Bas: f("bas"),
		Lym: f("lym"), Mon: f("mon"), Soe: f("soe"),
		Chol: f("chol"), Glu: f("glu"),
	}
	hb, err := services.PredictHba1c(hreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// LDL
	lreq := models.LdlRequest{Uid: q.Get("uid"), Age: i("age"), Gender: i("gender"),
		Rdw: f("rdw"), Wbc: f("wbc"), Rbc: f("rbc"),
		Hgb: f("hgb"), Hct: f("hct"), Mcv: f("mcv"),
		Mch: f("mch"), Mchc: f("mchc"), Plt: f("plt"),
		Neu: f("neu"), Eos: f("eos"), Bas: f("bas"),
		Lym: f("lym"), Mon: f("mon"), Soe: f("soe"),
		Chol: f("chol"), Glu: f("glu"),
	}
	ld, err := services.PredictLdl(lreq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// HDL
	hdlReq := models.HdlRequest{Uid: q.Get("uid"), Age: i("age"), Gender: i("gender"),
		Rdw: f("rdw"), Wbc: f("wbc"), Rbc: f("rbc"),
		Hgb: f("hgb"), Hct: f("hct"), Mcv: f("mcv"),
		Mch: f("mch"), Mchc: f("mchc"), Plt: f("plt"),
		Neu: f("neu"), Eos: f("eos"), Bas: f("bas"),
		Lym: f("lym"), Mon: f("mon"), Soe: f("soe"),
		Chol: f("chol"), Glu: f("glu"),
	}
	hd, err := services.PredictHdl(hdlReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TG
	treg := models.TgRequest{Uid: q.Get("uid"), Age: i("age"), Gender: i("gender"),
		Rdw: f("rdw"), Wbc: f("wbc"), Rbc: f("rbc"),
		Hgb: f("hgb"), Hct: f("hct"), Mcv: f("mcv"),
		Mch: f("mch"), Mchc: f("mchc"), Plt: f("plt"),
		Neu: f("neu"), Eos: f("eos"), Bas: f("bas"),
		Lym: f("lym"), Mon: f("mon"), Soe: f("soe"),
		Chol: f("chol"), Glu: f("glu"),
	}
	tg, err := services.PredictTg(treg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// FERR
	ferrReq := models.FerrRequest{Uid: q.Get("uid"), Age: i("age"), Gender: i("gender"),
		Rdw: f("rdw"), Wbc: f("wbc"), Rbc: f("rbc"),
		Hgb: f("hgb"), Hct: f("hct"), Mcv: f("mcv"),
		Mch: f("mch"), Mchc: f("mchc"), Plt: f("plt"),
		Neu: f("neu"), Eos: f("eos"), Bas: f("bas"),
		Lym: f("lym"), Mon: f("mon"), Soe: f("soe"),
		Crp: f("crp"),
	}
	fr, err := services.PredictFerr(ferrReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Собираем всё в один JSON
	combined := map[string]json.RawMessage{
		"hba1c": hb,
		"ldl":   ld,
		"hdl":   hd,
		"tg":    tg,
		"ferr":  fr,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(combined)
}
