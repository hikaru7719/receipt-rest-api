package model

import (
	"errors"
	"time"
)

// Receipt - レシートを表す構造体
type Receipt struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price" binding:"required"`
	Kind  string `json:"kind" binding:"required"`
	Date  string `json:"date" binding:"required"`
	Memo  string `json:"memo" binding:"required"`
}

// NewReceipt - レシートの生成
func NewReceipt(name string, price int, kind, date, memo string) (receipt *Receipt, err error) {
	if err := nameCheck(name); err != nil {
		return &Receipt{}, err
	}

	if err := priceCheck(price); err != nil {
		return &Receipt{}, err
	}

	if err := kindCheck(kind); err != nil {
		return &Receipt{}, err
	}

	if err := dateCheck(date); err != nil {
		return &Receipt{}, err
	}

	receipt = &Receipt{Name: name, Price: price, Kind: kind, Date: date, Memo: memo}

	return
}

func nameCheck(name string) (err error) {

	if len(name) > 20 {
		err = errors.New("invalid name")
	}

	return
}

func priceCheck(price int) (err error) {
	if price < 0 {
		err = errors.New("invalid price")
	}

	return
}

func kindCheck(kind string) (err error) {
	for _, value := range KIND {
		if kind == value {
			return
		}
	}

	err = errors.New("invalid type")
	return
}

func dateCheck(date string) (err error) {
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return err
	}

	return
}
