package models

import "gorm.io/gorm"

type Worker struct {
	Id          uint   `json:"id"`
	JobDesk     string `json:"jobdesk"`
	Domisili    string `json:"domisili"`
	TempatKerja string `json:"tempatkerja"`
	Desc        string `json:"desc"`
	Instagram   string `json:"instagram"`
	Github      string `json:"github"`
	Linkedin    string `json:"linkedin"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId"`
}

func (worker *Worker) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Worker{}).Count(&total)

	return total
}

func (worker *Worker) Take(db *gorm.DB, limit int, offset int) interface{} {
	var workers []Worker

	db.Preload("User").Offset(offset).Limit(limit).Find(&workers)

	return workers
}
