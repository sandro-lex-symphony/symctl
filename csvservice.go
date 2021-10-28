package symctl

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSV(file string, data [][]string) {
	csvFile, err := os.Create(file)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)
	for _, row := range data {
		_ = csvwriter.Write(row)
	}

	csvwriter.Flush()
	csvFile.Close()
}

func ShowCSV(data [][]string) {
	csvwriter := csv.NewWriter(os.Stdout)
	for _, row := range data {
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()

}
