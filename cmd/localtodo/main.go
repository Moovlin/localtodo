package main;

import (
  "fmt"
  "github.com/moovlin/localtodo/pkg/connection"
  "github.com/moovlin/localtodo/pkg/models"
  "github.com/moovlin/localtodo/pkg/datamodel"
)

func main(){
  fmt.Println("starting test")
  var test_db = "/home/sinn3r/.localtodo/todo.db"
  connector.Connect(test_db);
  connector.Migrate()

  var db = connector.Db


  var path = "/home/sinn3r/Downloads";
  var insDir = new(models.Directory)
  insDir.Path = path

  var todos = make([]models.Todo, 1)
  todos[0].TodoLevel = 0
  todos[0].Text = "Test"
  todos[0].Completed = false

  insDir.Todos = todos
  var affected, err = db.Insert(insDir)
  fmt.Println(affected, err)



  var dir models.Directory
  dir.Path = path
  var has, _ = db.Get(&dir)
  fmt.Println(has)
  connector.CloseConn()

  datamodel.OpenConnection(test_db)
  datamodel.GetTodosInDirectory(dir)
  datamodel.GetAllTodos()
  datamodel.CloseConnection()
}
