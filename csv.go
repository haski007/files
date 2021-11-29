package files

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

func WriteCSV(data interface{}, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file %s err: %w", filename, err)
	}

	if err := gocsv.MarshalFile(data, f); err != nil {
		return fmt.Errorf("gocsv MarshalFile err: %w", err)
	}
	return nil
}

func ReadCSV(path string, dest interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open file err: %w", err)
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &dest); err != nil {
		return fmt.Errorf("unmarshall csv err: %w", err)
	}
	return nil
}
