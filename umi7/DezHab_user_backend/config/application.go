package config

import (
	"encoding/json"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/umi7/DezHab_user_backend/dto"
	"gitlab.com/umi7/DezHab_user_backend/utils"
	"io/ioutil"
	"os"
)

const (
	configPath = "/src/gitlab.com/umi7/DezHab_user_backend/resources/"
)

var (
	configFile = "config.json"
)

var AppConfig dto.Config

func init() {
	fmt.Println(utils.GoPath)
	path := utils.GoPath + configPath + configFile
	fmt.Println(path)

	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	if err != nil {
		panic(err)
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, &AppConfig)
	if err != nil {
		panic(err)
	}
}
