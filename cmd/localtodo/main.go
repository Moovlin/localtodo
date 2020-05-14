package main;

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "github.com/moovlin/localtodo/pkg/connector"
)

func main(){
  var test_db = "/home/sinn3r/.localtodo/db.sqlite"
  conn = new(DBConnection{});
  conn.connect();

  var path = "/home/sinn3r";
  conn.Create(&Directory{Path: path});
  var dir Directory;
  db.First(&dir, "directory = ?", path);

  fmt.println(dir.Directory);
}
