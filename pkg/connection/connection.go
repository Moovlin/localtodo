package connector;

import (
  "fmt"
  _ "github.com/mattn/go-sqlite3"
      "xorm.io/xorm"
  "github.com/moovlin/localtodo/pkg/models"
)
var (
  Db *xorm.Engine
  err error
  connected bool
)


func Connect(db_path string) {
  if connected == false{
    Db, err = xorm.NewEngine("sqlite3", db_path);
    connected = true;
  }
  if err != nil {
    fmt.Println(err.Error())
    panic("failed to connect to db");
  }
  return
}

func CloseConn(){
  Db.Close()
  connected = false;
}

func Migrate(){
  Db.Sync2(new(models.Directory))
  Db.Sync2(new(models.Todo))
}
