package pkg

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func ParseCSV(filePath string) ([]Bank, error) {

	file,err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("filed to open file: %v", err)

	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV records: %v", err)
	}

	var banks []Bank
	for _, record := range records[1:] {
		bank := Bank{
			CountryISO2:  strings.ToUpper(record[0]),
			SwiftCode:    strings.TrimSpace(record[1]),
			CodeType:     strings.TrimSpace(record[2]),
			BankName:     strings.TrimSpace(record[3]),
			Address:      strings.TrimSpace(record[4]),
			TownName:     strings.TrimSpace(record[5]),
			CountryName:  strings.ToUpper(record[6]),
			TimeZone:     strings.TrimSpace(record[7]), 
			IsHeadquarter: false, 
		}

	if strings.HasSuffix(bank.SwiftCode, "XXX") {
		bank.IsHeadquarter = true
	}

	banks = append(banks, bank)
	}

	return banks, nil

}