package config

import (
	"encoding/json"
	"fmt"       //error printing
	"io/ioutil" //read config.json
)

var (
	//Token the bot token
	Token string
	//BotPrefix prefix used for the bot commands
	BotPrefix string
	config    *configStruct
)

type configStruct struct {
	Token     string `json:"token"`
	BotPrefix string `json:"botPrefix"`
}

//ReadConfig reads and loads data from configuration file
func ReadConfig() error {
	fmt.Println("Reading config file.")
	file, err := ioutil.ReadFile("./config.json") //reads file into 'file'. if an error occurs, places error in 'err'

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//print value of file by converting to string
	fmt.Println(string(file))
	//copying json values into config variable. returns error if one occurs.
	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	//store values in variables
	Token = config.Token
	BotPrefix = config.BotPrefix

	//no errors occured
	return nil
}
