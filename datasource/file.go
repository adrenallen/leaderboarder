package datasource

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type File struct {
	FilePath string
	entries  []Entry
}

func (f File) GetAll() []Entry {

	entries := []Entry{}
	file := f.getFile()
	reader := csv.NewReader(file)
	reader.Comma = f.getDelimiter()

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		scoreVal, _ := strconv.ParseFloat(record[1], 64)
		entries = append(entries, Entry{
			Name:  record[0],
			Score: scoreVal,
			Meta:  record[2],
		})
	}

	return entries

}

func (f File) SaveNew(entry Entry) {
	file := f.getFile()
	file.Seek(0, 2)
	csvWriter := csv.NewWriter(file)
	csvWriter.Comma = f.getDelimiter()

	//This is how to do it? Rob Pike pls
	scoreString := fmt.Sprintf("%f", entry.Score)

	err := csvWriter.WriteAll([][]string{
		{
			f.cleanStringDelimiters(entry.Name),
			scoreString,
			f.cleanStringDelimiters(entry.Meta),
		},
	})

	if err != nil {
		panic(err)
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}
}

func (f File) cleanStringDelimiters(val string) string {
	return strings.ReplaceAll(val, "|", "")
}

func (f File) getDelimiter() rune {
	return '|'
}

func (f File) getFile() *os.File {
	f.ensureFileExists()
	file, err := os.OpenFile(f.FilePath, os.O_RDWR, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	return file
}

func (f File) ensureFileExists() {
	_, err := os.Stat(f.FilePath)

	if os.IsNotExist(err) {
		print("Creating file at ", f.FilePath)
		f, err := os.Create(f.FilePath)
		if err != nil {
			panic(err)
		}
		f.Close()
	} else if err != nil {
		panic(err)
	}
}
