package files

import (
	"fmt"
	"os"
	"project/storage_accounts/storage_accounts/output"
)

type JsonDB struct {
	filename string
}

func NewJsonDB(name string) *JsonDB {
	return &JsonDB{
		filename: name,
	}
}

func (db *JsonDB) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *JsonDB) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		output.PrintError(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		output.PrintError(err)
		return
	}
	fmt.Println("Запись успешна")
}
