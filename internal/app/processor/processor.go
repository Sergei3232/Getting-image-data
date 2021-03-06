package processor

import (
	"github.com/Sergei3232/Getting-image-data/config"
	counter_id "github.com/Sergei3232/Getting-image-data/internal/app/counter"
	csvCounter "github.com/Sergei3232/Getting-image-data/internal/app/csv"
	"github.com/Sergei3232/Getting-image-data/internal/app/datastruct"
	"github.com/Sergei3232/Getting-image-data/internal/app/db"
	"log"
	"strconv"
	"sync"
)

const (
	portion     = 3000
	fileSaveCsv = "files/answer_files/test.csv"
)

type Processor struct {
	Counter   counter_id.TextCounter
	DbFileStr db.Repository
	DbImage   db.Repository
	CsvClient *csvCounter.ClientCsv
	Config    *config.WorkFile
}

func NewProcessor(config *config.WorkFile) (*Processor, error) {
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
		Config:    config,
	}, nil
}

func (p *Processor) Run() {
	log.Println("START SCRIPT")

	listCsvData, err := p.CsvClient.ReadWorkCsv(p.Config.PathWorkFile, p.Counter.GetIndent())
	lenArray := strconv.Itoa(len(listCsvData))

	if err != nil {
		log.Panicln(err)
	}
	var finishDataCsv = make([]datastruct.ImageFileCSV, 0)
	startPosition, endPosition := 0, portion
	for startPosition < len(listCsvData) {
		var countFile int
		var listForProcessing []datastruct.ImageFileCSV
		if endPosition < len(listCsvData) {
			countFile = portion
			listForProcessing = listCsvData[startPosition:endPosition]
		} else {
			listForProcessing = listCsvData[startPosition:]
			countFile = len(listForProcessing)
		}

		p.portionHandling(listForProcessing)
		finishDataCsv = append(finishDataCsv, listForProcessing...)
		p.CsvClient.WriterCsvFile(finishDataCsv, fileSaveCsv)

		lenFinisData := strconv.Itoa(len(finishDataCsv))
		log.Println("SCRIPT: " + lenFinisData + "/" + lenArray)

		startPosition, endPosition = startPosition+portion, endPosition+portion
		err = p.Counter.AddCountFile(uint64(countFile))
		if err != nil {
			log.Panicln(err)
		}
	}

	log.Println("END SCRIPT")
}

func (p *Processor) portionHandling(arrayPortion []datastruct.ImageFileCSV) {
	var wg sync.WaitGroup
	var countGorn int

	if len(arrayPortion)%500 == 0 {
		countGorn = len(arrayPortion) / 500
	} else {
		countGorn = len(arrayPortion)/500 + 1
	}

	wg.Add(countGorn)
	i, n := 0, 500
	for i < len(arrayPortion) {

		arrT := make([]datastruct.ImageFileCSV, 0, 500)
		if n > len(arrayPortion) {
			arrT = arrayPortion[i:]
		} else {
			arrT = arrayPortion[i:n]
		}

		go func() {
			defer wg.Done()
			mapCSVData := make(map[int64]datastruct.DataSCV)

			for n, val := range arrT {
				mapCSVData[val.MapiItem] = datastruct.DataSCV{Id: int64(n)}
			}

			p.DbImage.GettingIdImageFileStorage(mapCSVData)
			p.DbFileStr.GetImageHightWidth(mapCSVData)

			for _, val := range mapCSVData {
				arrT[val.Id].Height = val.Height
				arrT[val.Id].Width = val.Width
			}
		}()
		i += 500
		n += 500
	}

	wg.Wait()
}

func (p *Processor) CombiningFinishFiles() error {
	combinedDat, err := p.CsvClient.ReadMergingFiles(p.Config.PathFinishFile, 0)
	if err != nil {
		return err
	}
	p.CsvClient.WriterCsvFile(combinedDat, p.Config.PathFinishFile+"/finish.csv")

	return nil
}
