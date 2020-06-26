package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/BurntSushi/toml"
)

// StudentConfig 解析toml配置文件
type StudentConfig struct {
	Name     string `toml:"name"`
	ID       int    `toml:"id"`
	Password string `toml:"password"`
	Province string `toml:"province"`
	City     string `toml:"city"`
	District string `toml:"district"`
	Address  string `toml:"address"`
	Ymtys    int    `toml:"ymtys"`
	Tw       int    `toml:"tw"`
	Sfzx     int    `toml:"sfzx"`
	Sfcyglq  int    `toml:"sfcyglq"`
	Sfyzz    int    `toml:"sfyzz"`
	Cookie   string `toml:"cookie"`
	Path     string `toml:"path"`
}

// CollectConfigs 收集toml配置文件
func CollectConfigs(configDirectoryPath string) (studentConfigSlice []StudentConfig) {
	allFiles, err := ioutil.ReadDir(configDirectoryPath)
	if err != nil {
		fmt.Println("ioutil.ReadDir has error.", err)
	}

	for _, eachFileName := range allFiles {

		filenameWithSuffix := path.Base(eachFileName.Name())
		//filenameWithSuffix: 文件名带后缀。
		// fmt.Println("filenameWithSuffix =", filenameWithSuffix)

		fileSuffix := path.Ext(filenameWithSuffix)
		//fileSuffix: 文件后缀
		// fmt.Println("fileSuffix =", fileSuffix)
		if fileSuffix == ".toml" {
			configPath := fmt.Sprintf("%s/%s", configDirectoryPath, eachFileName.Name())
			studentConfigSlice = append(studentConfigSlice, ReadConfig(configPath))
		}
	}

	return studentConfigSlice
}

// ReadConfig 读取配置文件
func ReadConfig(configPath string) (tempConfig StudentConfig) {
	if _, err := toml.DecodeFile(configPath, &tempConfig); err != nil {
		log.Fatalln(err)
	}
	tempConfig.Path = configPath

	return
}

// UpdateConfig 更新配置文件
func UpdateConfig(newConfig StudentConfig) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(newConfig); err != nil {
		log.Fatal(err)
	}

	if ioutil.WriteFile(newConfig.Path, buf.Bytes(), 0644) == nil {
		fmt.Println("写入文件成功:", newConfig.Path)
	}
}
