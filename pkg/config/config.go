package config

import (
  "encoding/json"
  "io/ioutil"
  "fmt"
)

var(
  ConfPath string
  Conf ConfStruct
)

type ConfStruct struct {
  CollapseAll bool
  UncolHighlight bool
  DBPath string
}

func LoadConfig(){
  var confJson []byte
  var err error
  confJson, err = ioutil.ReadFile(ConfPath)
  if err != nil{
    fmt.Println(err.Error())
    panic("Error opening config")
  }
  json.Unmarshal([]byte(confJson), &Conf);
}

func SaveConfig(){
  data, _ := json.Marshal(Conf)
  var writeData = []byte(data)
  var err = ioutil.WriteFile(ConfPath, writeData, 0644)
  if err != nil{
    fmt.Println(err.Error())
    panic("Error writing conf to disk")
  }
}

func SaveDefaultConfig(){
  Conf = ConfStruct{CollapseAll: true, UncolHighlight: true, DBPath: "/home/sinn3r/.localtodo/todo.db"}
  SaveConfig()
}
