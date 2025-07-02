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

func WriterFile[T any](fileName string, objMap map[int]T) error {
	values := make([]T, 0, len(objMap))
	for _, v := range objMap {
		values = append(values, v)
	}

	file, err := os.OpenFile("docs/db/"+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("[\n"); err != nil {
		return err
	}

	for i, v := range values {
		b, err := json.Marshal(v)
		if err != nil {
			return err
		}

		if i < len(values)-1 {
			b = append(b, ',')
		}
		b = append(b, '\n')

		if _, err := file.Write(b); err != nil {
			return err
		}
	}

	if _, err := file.WriteString("]\n"); err != nil {
		return err
	}

	return nil
}