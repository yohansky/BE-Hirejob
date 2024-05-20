package models

import "gorm.io/gorm"

type Skill struct {
	Id       uint   `json:"id"`
	Nama     string `json:"nama"`
	WorkerId uint   //untuk relasi dengan worker
}

//relasi ke worker
// many2many ke worker

func (skill *Skill) Count(db *gorm.DB) int64 {
	var total int64

	db.Model(&Skill{}).Count(&total)

	return total
}

func (skill *Skill) Take(db *gorm.DB, limit int, offset int) interface{} {
	var skills []Skill

	db.Preload("User").Offset(offset).Limit(limit).Find(&skills)

	return skills
}
