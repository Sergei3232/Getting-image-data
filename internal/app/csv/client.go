package csvCounter

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Sergei3232/Getting-image-data/internal/app/datastruct"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type HandlersCsv interface {
	ReaderCsv()
	WriterCsv(listData []datastruct.ImageFileCSV, headerCsv []string) error
	WriteWorkCsv(pathFiles string) ([]datastruct.ImageFileCSV, error)
	GetListFilesProcess(path string) ([]string, error)
}

type ClientCsv struct {
	FilePath string
}

func NewCounterCsv() *ClientCsv {
	CounterCsv := &ClientCsv{}
	CounterCsv.ReaderCsv()
	return CounterCsv
}

func (c *ClientCsv) ReaderCsv() {

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

func (c *ClientCsv) GetListFilesProcess(path string) ([]string, error) {
	lst, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	files := make([]string, 0, len(lst))

	for _, val := range lst {
		if val.IsDir() {
			fmt.Printf("[%s]\n", val.Name())
		} else {
			files = append(files, path+"/"+val.Name())
		}
	}

	if len(files) == 0 {
		return files, errors.New("No files to process!")
	}
	return files, nil
}

func (c *ClientCsv) WriteWorkCsv(pathFiles string, countLineRead int) ([]datastruct.ImageFileCSV, error) {

	files, err := c.GetListFilesProcess(pathFiles)
	if err != nil {
		return nil, nil
	}
	dataCSV := make([]datastruct.ImageFileCSV, 0)
	for _, file := range files {
		f, err := os.Open(file)
		defer f.Close()

		if err != nil {
			return nil, err
		}

		lines, err := csv.NewReader(f).ReadAll()
		if err != nil {
			return nil, err
		}

		for nl, line := range lines { //line
			if nl > countLineRead {
				a1, _ := strconv.ParseInt(line[0], 10, 64)
				a2, _ := strconv.ParseInt(line[1], 10, 64)
				dataCSV = append(dataCSV, datastruct.ImageFileCSV{Sku: a1, MapiItem: a2})
			}
			nl++
		}
	}

	return dataCSV, nil
}
