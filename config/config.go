package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
)

type Server struct {
	Address string `json:"address"`
}

type Mongo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Database string `json:"database"`
	Timeout  int    `json:"timeout"`
}

type Config struct {
	Server `json:"server"`
	Key    string `json:"key"`
	Expire int    `json:"expire"`
	Mongo  `json:"mongo"`
}

var (
	Conf Config
)

func Init() {
	Conf = loadConfig()
}

func GetString(key string) string {
	return getKey(key).(string)
}

func GetInt(key string) int {
	return getKey(key).(int)
}

func GetObject(key string) interface{} {
	return getKey(key)
}

func getKey(key string) interface{} {
	if key == "" {
		return ""
	}
	lcaseKey := strings.ToLower(key)
	path := strings.Split(lcaseKey, ".")

	return findByMap(path, Conf)
}

func findByMap(path []string, source interface{}) interface{} {
	r := reflect.ValueOf(source)
	next := reflect.Indirect(r).FieldByName(strings.Title(path[0]))
	if len(path) > 1 {
		return findByMap(path[1:], next.Interface())
	}

	return next.Interface()
}

func loadConfig() Config {
	conf := Config{}
	file, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}
	err = json.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("[loadConfig]: %s\n", err)
	}

	return conf
}
