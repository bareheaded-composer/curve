package dao

import (
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"math/rand"
)

type FileStorage struct {
	rootPath string
}

func NewFileStorage(rootPath string) *FileStorage {
	return &FileStorage{
		rootPath: rootPath,
	}
}

func (s *FileStorage) StoreBase64Data(dirName string, storeName string, base64Data string) error {
	fileData, err := base64.StdEncoding.DecodeString(base64Data)
	if err != nil {
		logs.Error(err)
		return err
	}
	if err := s.Store(dirName, storeName, fileData); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (s *FileStorage) Store(dirName string, storeName string, data []byte) error {
	storePath := fmt.Sprintf("%s/%s/%s", s.rootPath, dirName, storeName)
	if err := ioutil.WriteFile(storePath, data, 0777); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

func (s *FileStorage) Get(dirName string, fileName string) ([]byte, error) {
	storePath := fmt.Sprintf("%s/%s/%s", s.rootPath, dirName, fileName)
	data, err := ioutil.ReadFile(storePath)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return data, nil
}

func (s *FileStorage) RandomGet(count int, dirName string) ([][]byte, error) {
	dirPath := fmt.Sprintf("%s/%s", s.rootPath, dirName)
	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if len(fileInfos) == 0 {
		return nil, nil
	}
	datas := make([][]byte, 0)
	for i := 0; i < count; i++ {
		fileIndex := rand.Intn(len(fileInfos))
		data, err := s.Get(dirName, fileInfos[fileIndex].Name())
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		datas = append(datas, data)
	}
	return datas, nil
}
