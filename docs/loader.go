package docs

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func ReadFileToMap[T any](fileName string) map[int]T {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	var elements []T
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, &elements); err != nil {
		panic(err)
	}

	result := make(map[int]T)
	for _, e := range elements {
		val := reflect.ValueOf(e)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		idField := val.FieldByName("ID")
		if !idField.IsValid() || idField.Kind() != reflect.Int {
			panic("type T must have a field `ID int`")
		}

		id := int(idField.Int())
		result[id] = e
	}

	return result
}

func WriterFile[T any](fileName string, obj T) error {
	b, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	b = append(b, '\n')

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(b)
	if err != nil {
		return err
	}

	return nil
}
