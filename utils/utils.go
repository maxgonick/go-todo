package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

var CfgDirPath string
var CfgFilePath string

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func InitConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	CfgDirPath = filepath.Join(home, ".config", "todo")
	if !pathExists(CfgDirPath) {
		os.Mkdir(CfgDirPath, 0777)
	}
	os.Chdir(CfgDirPath)

	CfgFilePath = filepath.Join(CfgDirPath, "config.json")
	if !pathExists(CfgFilePath) {
		f, err := os.Create("config.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
}

func MarshallToJSON(cfgFilePath string) TodoList {
	configData, err := os.ReadFile(cfgFilePath)
	if err != nil {
		panic(err)
	}

	var todoList TodoList

	if len(configData) == 0 {
		todoList = TodoList{
			Elements: []TodoElement{},
			NextId:   1,
		}
	} else {
		if err := json.Unmarshal(configData, &todoList); err != nil {
			panic(err)
		}
	}
	return todoList
}

func TodoListToDisk(data TodoList) {
	marshalledData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(CfgFilePath, marshalledData, 0777); err != nil {
		panic(err)
	}

	fmt.Println("Successfully wrote data to disk")
}
