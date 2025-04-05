package chapter_03

import (
	"encoding/csv"
	"log"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData []Record

func ReadWriteCsv() {
	if len(os.Args) < 3 {
		log.Println("In- and output required")
		os.Exit(1)
	}
	input, output := os.Args[1], os.Args[2]
	lines, err := readCsvFile(input)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// Not safe. Assumes each line always has 4 entries!
	for _, line := range lines {
		temp := Record{
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp)
		log.Println(temp)
	}
	err = saveCsvFile(output)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func readCsvFile(filePath string) ([][]string, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	// Not optimal. Reads all lines at once. Could be a problem for huge files.
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return lines, nil
}

func saveCsvFile(filepath string) error {
	csvFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	// Not a normal operation - change separation by comma to tab separation
	csvWriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		err = csvWriter.Write(temp)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil
}
