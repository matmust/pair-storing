package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/matmust/pairStoring"
)

type FileStorage struct {
	repository pairStoring.PairRepository
	filename   string
}

// NewFileStorage returns a new  instance of FileStorage.
func NewFileStorage(filename string, repository pairStoring.PairRepository) *FileStorage {
	return &FileStorage{filename: filename, repository: repository}
}

// Load, loads data to in-memory database from storage file if exists.
func (fs *FileStorage) Load() error {

	f, err := os.OpenFile(fs.filename, os.O_CREATE, os.ModePerm)

	if err != nil {
		fmt.Println("could not open file error:", err.Error())
		return err
	}
	m := make(map[string]string)
	if err := json.NewDecoder(f).Decode(&m); err == io.EOF {
		return nil
	} else if err != nil {
		fmt.Println("could not decode file error:", err.Error())
		return err
	}

	fs.repository.SetAll(m)
	return nil
}

// Store, backups data to storage file from in-memory database.
func (fs *FileStorage) Store() error {

	f, err := os.Create(fs.filename)
	if err != nil {
		fmt.Println("could not open file error:", err.Error())
		return err
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(fs.repository.GetAll()); err != nil {
		fmt.Println("could not encode file error:", err.Error())
		return err
	}
	return nil

}

func (fs *FileStorage) PeriodicBackup(duration time.Duration) {
	ticker := time.NewTicker(duration)

	go func() {
		for range ticker.C {
			err := fs.Store()
			if err != nil {
				fmt.Println("Periodic backup error:", err)
			}
		}
	}()
}
