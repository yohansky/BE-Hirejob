package models

import "gorm.io/gorm"

type Worker struct {
	Id          uint   `json:"id"`
	Nama        string `json:"nama"`
	JobDesk     string `json:"jobdesk"`
	Domisili    string `json:"domisili"`
	TempatKerja string `json:"tempatkerja"`
	Desc        string `json:"desc"`
	Instagram   string `json:"instagram"`
	Github      string `json:"github"`
	Linkedin    string `json:"linkedin"`
	UserId      uint
	User        User `gorm:"foreignKey:UserId"`
	// SkillId     uint
	Skill []Skill `gorm:"foreignKey:WorkerId"` // kolom WorkerId ada di struct skill
	// ProjectId    uint
	// Project      Project `gorm:"foreignKey:ProjectId"`
	// ExperienceId uint
	// Experience   Experience `gorm:"foreignKey:ExperienceId"`
}

// many2many dengan skill
//relasi ke skill,

func (worker *Worker) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Worker{}).Count(&total)

	return total
}

func (worker *Worker) Take(db *gorm.DB, limit int, offset int) interface{} {
	var workers []Worker

	db.Preload("User").Preload("Skill").Offset(offset).Limit(limit).Find(&workers)

	return workers
}
