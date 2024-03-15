package CSV

import (
	"encoding/csv"
	"log"
	"os"
)

type CSV struct {
	Records [][]string
	Headers []string
}

func GetCsvHeaders(records [][]string) []string {
	return records[0]
}

func ReadCsvFile(filePath string) CSV {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return CSV{records, GetCsvHeaders(records)}
}
