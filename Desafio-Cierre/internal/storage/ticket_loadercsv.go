package loader

import (
	"encoding/csv"
	"fmt"
	"goweb/Desafio-Cierre/internal"
	"io"
	"os"
	"strconv"
)

// NewLoaderTicketCSV creates a new ticket loader from a CSV file
func NewLoaderTicketCSV(filePath string) *LoaderTicketCSV {
	return &LoaderTicketCSV{
		filePath: filePath,
	}
}

// LoaderTicketCSV represents a ticket loader from a CSV file
type LoaderTicketCSV struct {
	filePath string
}

// Load loads the tickets from the CSV file
func (l *LoaderTicketCSV) Load() (t map[int]internal.TicketAttributes, lastId int, err error) {
	// open the file
	f, err := os.Open(l.filePath)
	if err != nil {
		err = fmt.Errorf("error opening file: %v", err)
		return
	}
	defer f.Close()

	// read the file
	r := csv.NewReader(f)

	// read the records
	t = make(map[int]internal.TicketAttributes)
	for {
		record, ok := r.Read()
		if ok != nil {
			if ok == io.EOF {
				break
			}

			err = fmt.Errorf("error reading record: %v", ok)
			return
		}

		// serialize the record
		id, ok := strconv.Atoi(record[0])
		price, ok := strconv.ParseFloat(record[5], 64)
		ticket := internal.TicketAttributes{
			Name:    record[1],
			Email:   record[2],
			Country: record[3],
			Hour:    record[4],
			Price:   price,
		}

		// add the ticket to the map
		t[id] = ticket
		lastId = id
	}

	return
}
