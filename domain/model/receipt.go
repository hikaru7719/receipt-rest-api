package model

import (
	"errors"
	"time"
)

type Receipt struct {
	ID   int
	Name string
	Kind string
	Date string
	Memo string
}

func NewReceipt(name, kind, date, memo string) (receipt Receipt, err error) {
	if err := nameCheck(name); err != nil {
		return Receipt{}, err
	}

	if err := kindCheck(kind); err != nil {
		return Receipt{}, err
	}

	if err := dateCheck(date); err != nil {
		return Receipt{}, err
	}

	receipt = Receipt{Name: name, Kind: kind, Date: date, Memo: memo}

	return
}

func nameCheck(name string) (err error) {

	if len(name) > 20 {
		err = errors.New("invalid name")
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
