package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/BurntSushi/toml"
)

// StudentConfig 解析toml配置文件
type StudentConfig struct {
	Name              string    `toml:"name"`
	ID                int       `toml:"id"`
	Password          string    `toml:"password"`
	Province          string    `toml:"province"`
	City              string    `toml:"city"`
	Area              string    `toml:"area"`
	Address           string    `toml:"address"`
	Szgjcs            string    `toml:"szgjcs"`
	Szcs              string    `toml:"szcs"`
	Szgj              string    `toml:"szgj"`
	Zgfxdq            string    `toml:"zgfxdq"` // 中高风险地区
	Mjry              string    `toml:"mjry"`   // 密接人员
	Csmjry            string    `toml:"csmjry"`
	Tw                string    `toml:"tw"`      // 体温
	Sfzx              string    `toml:"sfzx"`    // 是否在校
	Sfcyglq           string    `toml:"sfcyglq"` // 是否处于隔离期
	Sfcxtz            string    `toml:"sfcxtz"`
	Sfjcbh            string    `toml:"sfjcbh"`
	Sfcxzysx          string    `toml:"sfcxzysx"`
	Qksm              string    `toml:"qksm"`
	Sfyyjc            string    `toml:"sfyyjc"`
	Jcjgqr            string    `toml:"jcjgqr"`
	Remark            string    `toml:"remark"`
	Sfjcwhry          string    `toml:"sfjcwhry"` // 是否接触武汉人员
	Sfjchbry          string    `toml:"sfjchbry"` // 是否接触河北人员
	Gllx              string    `toml:"gllx"`
	Glksrq            string    `toml:"glksrq"`
	Jcbhlx            string    `toml:"jcbhlx"`
	Jcbhrq            string    `toml:"jcbhrq"`
	Ismoved           string    `toml:"ismoved"`
	Bztcyy            string    `toml:"bztcyy"`
	Sftjhb            string    `toml:"sftjhb"`   // 是否途径河北
	Sftjwh            string    `toml:"sftjwh"`   // 是否途径武汉
	Sfjcjwry          string    `toml:"sfjcjwry"` // 是否接触境外人员
	Jcjg              string    `toml:"jcjg"`
	SCKEY             string    `toml:"SCKEY"`
	Cookie            string    `toml:"cookie"`
	Path              string    `toml:"path"`
	LastestUpdateTime time.Time `toml:"lastestupdatetime"`
}

type MainConfig struct {
	Cron        string `toml:"cron"`
	BaseURL     string `toml:"BaseURL"`
	LoginURL    string `toml:"LoginURL"`
	SaveURL     string `toml:"SaveURL"`
	MyUserAgent string `toml:"MyUserAgent"`
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

		if filenameWithSuffix != "main.toml" {
			fileSuffix := path.Ext(filenameWithSuffix)
			//fileSuffix: 文件后缀
			// fmt.Println("fileSuffix =", fileSuffix)
			if fileSuffix == ".toml" {
				configPath := fmt.Sprintf("%s/%s", configDirectoryPath, eachFileName.Name())
				studentConfigSlice = append(studentConfigSlice, ReadConfig(configPath))
			}
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
	// fmt.Println("配置文件", tempConfig)

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

func ReadMainConfig(mainConfigPath string) (mainConfig MainConfig) {
	if _, err := toml.DecodeFile(mainConfigPath, &mainConfig); err != nil {
		log.Fatalln(err)
	}

	return
}
