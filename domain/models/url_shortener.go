package models

import (
	"github.com/ventu-io/go-shortid"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Link string `json:"url" gorm:"size:200;not null;"`
	Hash string `json:"hash" gorm:"size:10;not null;"`
}

// BeforeCreate func process data before create object
func (u *Url) BeforeCreate(tx *gorm.DB) (err error) {
	sid, _ := shortid.New(1, shortid.DefaultABC, 4242)
	hash, err := sid.Generate()
	if err != nil {
		return
	}
	u.Hash = hash
	return nil
}
