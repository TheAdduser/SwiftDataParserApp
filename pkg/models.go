package pkg

import "github.com/jinzhu/gorm"

type Bank struct {
	ID	 					uint   `json:"id" gorm:"primary_key"`
	SwiftCode 		string `json:"swift_code" gorm:"unique;not null"`
	Address 			string `json:"address"`
	BankName 			string `json:"bank_name"`
	CountryISO2 	string `json:"country_iso2"`
	CountryName 	string `json:"country_name"`
	IsHeadquarter bool `json:"is_headquarter"`
	HedquaterID 	*uint `json:"headquarter_id"`
}