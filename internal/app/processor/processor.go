package processor

import (
	counter_id "github.com/Sergei3232/Getting-image-data/internal/app/counter"
	csvCounter "github.com/Sergei3232/Getting-image-data/internal/app/csv"
	"github.com/Sergei3232/Getting-image-data/internal/app/db"
)

type Processor struct {
	Counter   counter_id.TextCounter
	DbFileStr db.Repository
	DbImage   db.Repository
	CsvClient csvCounter.HandlersCsv
}
