package counter_id

import (
	"io/ioutil"
	"strconv"
)

type TextCounter interface {
	AddCountFile(id int) error
	readFileCount() error
}

type TextCounterStruct struct {
	FilePath string
	Indent   uint64
}

func NewCounter(filePath string, indent uint64) (*TextCounterStruct, error) {
	counter := &TextCounterStruct{filePath, indent}
	err := counter.readFileCount()
	if err != nil {
		return nil, err
	}

	return counter, nil
}

func (t *TextCounterStruct) AddCountFile(count uint64) error {
	t.Indent += count
	strId := strconv.Itoa(int(t.Indent))

	var d = []byte(strId)

	err := ioutil.WriteFile(t.FilePath, d, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (t *TextCounterStruct) readFileCount() error {
	f, err := ioutil.ReadFile(t.FilePath)
	if err != nil {
		return err
	}

	lastId, err := strconv.ParseUint(string(f), 10, 64)

	if err != nil {
		return err
	}
	t.Indent = lastId
	return nil
}
