package model

type Quote struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Code       string `gorm:"code"`
	Codein     string `gorm:"codein"`
	Name       string `gorm:"name"`
	High       string `gorm:"high"`
	Low        string `gorm:"low"`
	VarBid     string `gorm:"varBid"`
	PctChange  string `gorm:"pctChange"`
	Bid        string `gorm:"bid"`
	Ask        string `gorm:"ask"`
	Timestamp  string `gorm:"timestamp"`
	CreateDate string `gorm:"create_date"`
}
