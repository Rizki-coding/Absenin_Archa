package model

import (
	"time"

	"gorm.io/gorm"
)

type Presen struct {
	gorm.Model
	Id_pengguna string `json:"id_pengguna" gorm:"NOT NULL"`
	Time_in     time.Time
	Lokasi_in   string `json:"lokasi_in"`
	Time_out    time.Time
	Lokasi_out  string `json:"lokasi_out"`
}

type PresenJResponses struct {
	Type    bool     `json:"status"`
	Message string   `json:"message"`
	Data    []Presen `json:"data"`
}
