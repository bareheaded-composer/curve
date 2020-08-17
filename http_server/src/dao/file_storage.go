package dao

import (
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

type FileStorage struct {
	rootPath string
}

func NewFileStorage(rootPath string) *FileStorage {
	return &FileStorage{
		rootPath: rootPath,
	}
}

func (s *FileStorage) StoreBase64Data(storeName string, base64Data string) error {
	fileData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		logs.Error(err)
		return err
	}
	if err := s.Store(storeName, fileData); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (s *FileStorage) Store(storeName string, data []byte) error {
	storePath := fmt.Sprintf("%s/%s", s.rootPath, storeName)
	if err := ioutil.WriteFile(storePath, data, 0777); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (s *FileStorage) Get(fileName string) ([]byte, error) {
	storePath := fmt.Sprintf("%s/%s", s.rootPath, fileName)
	data, err := ioutil.ReadFile(storePath)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return data, nil
}
