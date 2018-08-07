package model

import (
	"testing"
)

func TestNameCheck(t *testing.T) {
	safeName := "aaaaaaaaaaaaaaaaaaaa"
	invalidName := "aaaaaaaaaaaaaaaaaaaaa"

	if err := nameCheck(safeName); err != nil {
		t.Error("lenght 20 is ok.")
	}

	if err := nameCheck(invalidName); err == nil {
		t.Error("lenght 21 is invalid.")
	}

}

func TestDateCheck(t *testing.T) {
	safeFormat := "2018-05-02"
	numberFormat := "1111111111"
	invalidMonth := "2018-33-02"
	if err := dateCheck(safeFormat); err != nil {
		t.Error(err)
	}

	if err := dateCheck(numberFormat); err == nil {
		t.Error("number is invalid format")
	}

	if err := dateCheck(invalidMonth); err == nil {
		t.Error("month 33 is invalid format")
	}

}

func TestKindCheck(t *testing.T) {
	safeKind := []string{"食費", "日用品", "趣味,娯楽", "交通費", "衣類,美容", "その他"}
	for _, value := range safeKind {
		if err := kindCheck(value); err != nil {
			t.Error("safe kind")
		}
	}

	invalidKind := "aaaa"
	if err := kindCheck(invalidKind); err == nil {
		t.Error("invalid kind")
	}
}

func TestNewReceipt(t *testing.T) {
	expected := Receipt{name: "キャップ", kind: "衣類,美容", date: "2018-08-07", memo: "キャップを購入した。"}
	actual, _ := NewReceipt("キャップ", "衣類,美容", "2018-08-07", "キャップを購入した。")
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
