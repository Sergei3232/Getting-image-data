package csvCounter

import (
	"encoding/csv"
	"github.com/Sergei3232/Getting-image-data/internal/app/datastruct"
	"log"
	"os"
	"strconv"
)

type HandlersCsv interface {
	ReaderCsv()
	WriterCsv() error
}

type ClientCsv struct {
	FilePath string
}

func NewCounterCsv(filePath string) *ClientCsv {
	CounterCsv := &ClientCsv{filePath}
	CounterCsv.ReaderCsv()
	return CounterCsv
}

func (c *ClientCsv) ReaderCsv() {
	//f, err := os.Open(file)
	//defer f.Close()
	//
	//if err != nil {
	//	return err
	//}
	//lines, err := csv.NewReader(f).ReadAll()
	//if err != nil {
	//	return err
	//}
	//
	//for nl, line := range lines {
	//	if nl > 0 {
	//		db_csv.InsertDataCSVFile(line, fileName)
	//		fmt.Println(line)
	//	}
	//	nl++
	//}
}

func (c *ClientCsv) WriterCsv(listData []datastruct.ImageFileCSV, headerCsv []string) error {
	records := [][]string{{}}
	records = append(records, headerCsv)

	for _, val := range listData {
		records = append(records, []string{
			strconv.FormatUint(uint64(val.Sku), 10),
			strconv.FormatUint(uint64(val.MapiItem), 10),
			strconv.FormatUint(uint64(val.Height), 10),
			strconv.FormatUint(uint64(val.Width), 10),
		})
	}

	file, errCreate := os.Create(c.FilePath)
	if errCreate != nil {
		log.Panic(errCreate)
	}

	w := csv.NewWriter(file)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	w.Flush()

	return nil
}
