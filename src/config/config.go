/*
Package config default config
*/
package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config 配置
type Config struct {
	AppInfo appInfo `json:"AppInfo"`
	Log     logConf `json:"Log"`
	Wechat  wechat  `json:"Wechat"`
	URL     url     `json:"URL"`
}

type appInfo struct {
	Env  string `json:"Env"` // example: local, dev, prod
	Addr string `json:"Addr"`
}

type logConf struct {
	LogBasePath string `json:"LogBasePath"`
	LogFileName string `json:"LogFileName"`
}

type wechat struct {
	AppID        string `json:"AppID"`
	AppSecret    string `json:"AppSecret"`
	RedirectURI  string `json:"RedirectURI"`
	ResponseType string `json:"ResponseType"`
}

type url struct {
	CodeURL     string `json:"CodeURL"`
	TokenURL    string `json:"TokenURL"`
	RefreshURL  string `json:"RefreshURL"`
	UserInfoURL string `json:"UserInfoURL"`
}

// Conf 配置
var Conf *Config

var filePrefix = "/app/src/config/"

func init() {
	log.Println("begin init configs")
	initConf()
	readEnv()
	log.Println("over init configs")
}

func initConf() {

	log.Println("begin init default config")

	Conf = &Config{}
	fileName := "default.json"

	// read default config
	data, err := ioutil.ReadFile(filePrefix + fileName)
	if err != nil {
		log.Println("config-initConf: read default.json error")
		log.Panic(err)
		return
	}
	err = json.Unmarshal(data, Conf)
	if err != nil {
		log.Println("config-initConf: unmarshal default.json error")
		log.Panic(err)
		return
	}
	log.Println("over init default config")
}

func readEnv() {
	log.Println("begin read env")
	if v, ok := os.LookupEnv("Env"); ok {
		Conf.AppInfo.Env = v
	}
	if v, ok := os.LookupEnv("WEIXIN_APPID"); ok {
		Conf.Wechat.AppID = v
	}

	if v, ok := os.LookupEnv("WEIXIN_APPSECRET"); ok {
		Conf.Wechat.AppSecret = v
	}

	if v, ok := os.LookupEnv("APP_ADDR"); ok {
		Conf.AppInfo.Addr = v
	}
	log.Println("over read env")
}
