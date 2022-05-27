package processor

import (
	"github.com/Sergei3232/Getting-image-data/config"
	counter_id "github.com/Sergei3232/Getting-image-data/internal/app/counter"
	csvCounter "github.com/Sergei3232/Getting-image-data/internal/app/csv"
	"github.com/Sergei3232/Getting-image-data/internal/app/db"
)

type Processor struct {
	Counter   counter_id.TextCounter
	DbFileStr db.Repository
	DbImage   db.Repository
	CsvClient *csvCounter.ClientCsv
}

func NewProcessor(config config.WorkFile) (*Processor, error) {
	counter, err := counter_id.NewCounter(config.PathCountFile, 0)
	if err != nil {
		return nil, err
	}
	csvCounter := csvCounter.NewCounterCsv()
	dbFileStr, err := db.NewDbConnectClient(config.DNSFileLoader)
	if err != nil {
		return nil, err
	}
	ddbImage, err := db.NewDbConnectClient(config.DNSImageLoader)
	if err != nil {
		return nil, err
	}
	return &Processor{
		Counter:   counter,
		CsvClient: csvCounter,
		DbImage:   ddbImage,
		DbFileStr: dbFileStr,
	}, nil
}
