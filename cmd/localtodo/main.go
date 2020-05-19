package main;

import (
  "fmt"
  "path/filepath"
  "github.com/moovlin/localtodo/pkg/connection"
  "github.com/moovlin/localtodo/pkg/dirinfo'"
  "github.com/moovlin/localtodo/pkg/models"
  // "github.com/moovlin/localtodo/pkg/datamodel"
  // "github.com/moovlin/localtodo/pkg/config"
)

func main(){
  fmt.Println("starting test")
  var test_db = "/home/sinn3r/.localtodo/todo.db"
  connector.Connect(test_db);
  connector.Migrate()

  // var db = connector.Db
  var testDir models.Directory{}
  testDir.Path = ""
  }
