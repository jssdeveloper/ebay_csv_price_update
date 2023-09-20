package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Update struct {
	Sku   string
	Price string
}

type Active struct {
	Action     string
	ItemNumber string
	Sku        string
	Price      string
}

var (
	update []Update
	active []Active
)

func main() {
	readCsv([]string{"update.csv", "active.csv"})
	updateMap := craeteMap()

	// Write to output csv
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Failed to craete output file")
	}
	defer file.Close()

	csv := csv.NewWriter(file)
	csv.Write([]string{"Action", "Item number", "Start price"})

	for _, v := range active {
		data, ok := updateMap[v.Sku]
		if ok {
			csv.Write([]string{"Revise", v.ItemNumber, data})
		}
	}

	csv.Flush()

	fmt.Println("Finished writing csv file")

}

// Reads csv files to structs
func readCsv(files []string) {
	for _, f := range files {
		file, err := os.Open(f)
		if err != nil {
			fmt.Printf("Failed to open Csv File %v\n", file)
			os.Exit(2)
		}
		defer file.Close()

		csv := csv.NewReader(file)

		// ebay de csv response has delimiter ';'
		if f == "active.csv" {
			csv.Comma = ';'
		}
		rows, err := csv.ReadAll()

		if f == "active.csv" {
			for _, v := range rows {
				active = append(active, Active{Action: "Revise", ItemNumber: v[0], Sku: v[3], Price: v[7]})
			}
		} else {
			for _, v := range rows {
				update = append(update, Update{Sku: v[0], Price: v[1]})
			}
		}
	}
}

func craeteMap() map[string]string {
	updateMap := map[string]string{}

	for _, v := range update {
		updateMap[v.Sku] = v.Price
	}

	return updateMap
}
