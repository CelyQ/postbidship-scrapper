package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func createExcelFile(tableMap map[string][]string) string {
	f := excelize.NewFile()

	f.SetCellValue("Sheet1", "A1", "State")
	f.SetCellValue("Sheet1", "B1", "Company")
	f.SetCellValue("Sheet1", "C1", "Load")
	f.SetCellValue("Sheet1", "D1", "Host #")
	f.SetCellValue("Sheet1", "E1", "Contract")
	f.SetCellValue("Sheet1", "F1", "Type")
	f.SetCellValue("Sheet1", "G1", "Source")
	f.SetCellValue("Sheet1", "H1", "Destination")
	f.SetCellValue("Sheet1", "I1", "Stops")
	f.SetCellValue("Sheet1", "J1", "Distance")
	f.SetCellValue("Sheet1", "K1", "Start")
	f.SetCellValue("Sheet1", "L1", "End")
	f.SetCellValue("Sheet1", "M1", "Weight")
	f.SetCellValue("Sheet1", "N1", "Equipment")
	f.SetCellValue("Sheet1", "O1", "Post Start")
	f.SetCellValue("Sheet1", "P1", "Post End")
	f.SetCellValue("Sheet1", "Q1", "Max Bid")
	f.SetCellValue("Sheet1", "R1", "Cargo Value")
	f.SetCellValue("Sheet1", "S1", "Post Type")
	f.SetCellValue("Sheet1", "T1", "Trailer Length")
	f.SetCellValue("Sheet1", "U1", "Freight Class")
	f.SetCellValue("Sheet1", "V1", "Reefer Min Temp")
	f.SetCellValue("Sheet1", "W1", "Reefer Max Temp")
	f.SetCellValue("Sheet1", "X1", "Temperature Scale")
	f.SetCellValue("Sheet1", "Y1", "Team Required")
	f.SetCellValue("Sheet1", "Z1", "Hazmat Required")
	f.SetCellValue("Sheet1", "AA1", "Bid State")
	f.SetCellValue("Sheet1", "AB1", "Lowest Bid")

	for i, value := range tableMap["State"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), value)
	}

	for i, value := range tableMap["Company"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), value)
	}

	for i, value := range tableMap["Load"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), value)
	}

	for i, value := range tableMap["Host #"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), value)
	}

	for i, value := range tableMap["Contract"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), value)
	}

	for i, value := range tableMap["Type"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), value)
	}

	for i, value := range tableMap["Source"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), value)
	}

	for i, value := range tableMap["Destination"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), value)
	}

	for i, value := range tableMap["Stops"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("I%d", i+2), value)
	}

	for i, value := range tableMap["Distance"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("J%d", i+2), value)
	}

	for i, value := range tableMap["Start"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("K%d", i+2), value)
	}

	for i, value := range tableMap["End"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("L%d", i+2), value)
	}

	for i, value := range tableMap["Weight"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("M%d", i+2), value)
	}

	for i, value := range tableMap["Equipment"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("N%d", i+2), value)
	}

	for i, value := range tableMap["Post Start"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("O%d", i+2), value)
	}

	for i, value := range tableMap["Post End"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("P%d", i+2), value)
	}

	for i, value := range tableMap["Max Bid"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%d", i+2), value)
	}

	for i, value := range tableMap["Cargo Value"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("R%d", i+2), value)
	}

	for i, value := range tableMap["Post Type"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("S%d", i+2), value)
	}

	for i, value := range tableMap["Trailer Length"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("T%d", i+2), value)
	}

	for i, value := range tableMap["Freight Class"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("U%d", i+2), value)
	}

	for i, value := range tableMap["Reefer Min Temp"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("V%d", i+2), value)
	}

	for i, value := range tableMap["Reefer Max Temp"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("W%d", i+2), value)
	}

	for i, value := range tableMap["Temperature Scale"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("X%d", i+2), value)
	}

	for i, value := range tableMap["Team Required"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("Y%d", i+2), value)
	}

	for i, value := range tableMap["Hazmat Required"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("Z%d", i+2), value)
	}

	for i, value := range tableMap["Bid State"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("AA%d", i+2), value)
	}

	for i, value := range tableMap["Lowest Bid"] {
		f.SetCellValue("Sheet1", fmt.Sprintf("AB%d", i+2), value)
	}

	fileName := "OldCastle.xlsx"
	f.SaveAs(fileName)

	return fileName
}
