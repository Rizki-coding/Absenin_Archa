package module

import (
	"APLIKASI_1/model"
	"encoding/json"
	"net/http"
	"time"
)

func GetPresen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var presen []model.Presen
	DB.Find(&presen)
	var response = model.PresenJResponses{Type: true, Message: "Success Get data", Data: presen}
	json.NewEncoder(w).Encode(response)
}

func Presen(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var presens []model.Presen
	var presen model.Presen
	var response = model.PresenJResponses{}

	Time := time.Now()
	Id_pengguna := r.FormValue("id_pengguna")

	result := DB.Where("id_pengguna = ? AND TO_CHAR(time_in, 'YYYY-MM-DD') = LEFT( ? , 10)", Id_pengguna, Time).First(&presens)

	if result.RowsAffected > 0 {
		DB.First(&presen, "id_pengguna = ?", r.FormValue("id_pengguna"))
		presen.Time_out = Time
		presen.Lokasi_out = r.FormValue("lokasi_out")
		DB.Save(&presen)
		response = model.PresenJResponses{Type: true, Message: "success, Check out"}
	} else {
		presen.Id_pengguna = r.FormValue("id_pengguna")
		presen.Time_in = Time
		presen.Lokasi_in = r.FormValue("lokasi_in")
		DB.Create(&presen)
		response = model.PresenJResponses{Type: true, Message: "Success, Check in"}
	}
	json.NewEncoder(w).Encode(response)
}
