package models

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNameIsRequired     = errors.New("name é obrigatório")
	ErrServiceIsRequired  = errors.New("service é obrigatório")
	ErrDeadlineIsRequired = errors.New("deadline é obrigatório")
	ErrPriceIsRequired    = errors.New("price é obrigatório")
)

func (Quote) TableName() string {
	return "quotes"
}

type Quote struct {
	ID       int    `gorm:"primaryKey;autoIncrement" column:"id" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Service  string `gorm:"column:service" json:"service"`
	Deadline string `gorm:"column:deadline" json:"deadline"`
	Price    int    `gorm:"column:price" json:"price"`
}

func (q *Quote) NewQuote(db *gorm.DB) (*Quote, error) {
	var err error
	err = db.Debug().Create(&q).Error
	if err != nil {
		return &Quote{}, err
	}
	return q, nil
}

func (q *Quote) GetQuotes(db *gorm.DB) (*[]Quote, error) {
	var err error
	quotes := []Quote{}
	err = db.Debug().Model(&Quote{}).Limit(100).Find(&quotes).Error
	if err != nil {
		return &[]Quote{}, err
	}
	return &quotes, err
}

func (q *Quote) Validate() error {
	if q.Name == "" {
		return ErrNameIsRequired
	}
	if q.Service == "" {
		return ErrServiceIsRequired
	}
	if q.Deadline == "" {
		return ErrDeadlineIsRequired
	}
	if q.Price == 0 {
		return ErrPriceIsRequired
	}
	return nil
}
