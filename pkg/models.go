package pkg

import "github.com/jinzhu/gorm"

type Bank struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	CountryISO2   string `json:"countryISO2"`
	SwiftCode     string `json:"swiftCode"`
	CodeType      string `json:"codeType"`
	BankName      string `json:"bankName"`
	Address       string `json:"address"`
	TownName      string `json:"townName"`
	CountryName   string `json:"countryName"`
	TimeZone      string `json:"timeZone"`
	IsHeadquarter bool   `json:"isHeadquarter"`
	HeadquarterID *uint  `json:"headquarterId,omitempty"`
}