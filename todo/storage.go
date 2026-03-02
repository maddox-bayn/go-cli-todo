package todo

import (
	"encoding/json"
	"os"
)

//"encoding/json"
//"os"

type TodoStorage interface {
	Load() ([]Todo, error)
	Save([]Todo) error
}

type FileStorage struct {
	dataFile string
}

func NewFileStorage(dataFile string) *FileStorage {
	return &FileStorage{dataFile: dataFile}
}
func (f *FileStorage) Load() ([]Todo, error) {
	data, err := os.ReadFile(f.dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, err
		}
		return []Todo{}, err
	}

	if len(data) == 0 {
		return []Todo{}, err
	}

	var todos []Todo
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}
	return todos, nil
}
func (f *FileStorage) Save(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.dataFile, data, 0644)
}
