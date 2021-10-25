package files

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJson - here everything is clear
func ReadJson(path string, dest interface{}) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file err: %s", err)
	}

	if err := json.Unmarshal(data, &dest); err != nil {
		return fmt.Errorf("json unmarshall file [%s] data err: %s", path, err)
	}

	return err
}

// WriteJson - here everything is clear
func WriteJson(file string, data interface{}) error {
	newData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("[json.MarshallIndent] err: %s", err)
	}

	return ioutil.WriteFile(file, newData, os.ModePerm)
}