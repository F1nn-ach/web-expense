package models

type Transaction struct {
	Date     string `json:"date"`
	Item     string `json:"item"`
	Type     string `json:"type"` // รับ/จ่าย
	Amount   int    `json:"amount"`
	Method   string `json:"method"` // สด/ธนาคาร
	Category string `json:"category"`
	Note     string `json:"note"`
}
