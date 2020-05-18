package datamodel

import (
  "fmt"
  "github.com/moovlin/localtodo/pkg/connection"
  "github.com/moovlin/localtodo/pkg/models"
)

func OpenConnection(path string){
  connector.Connect(path)
}

func GetTodosInDirectory(dir *models.Directory){
  var db = connector.Db

  var found, err = db.Get(dir)
  if err != nil{
    fmt.Println(err.Error())
    panic("Couldn't get the table")
  }
  if found == false{
    fmt.Println("Couldn't find it")
  }

  for _, todo := range dir.Todos{
    fmt.Println(todo.Text)
  }
}

func GetAllTodos(dirs []models.Directory){
  var db = connector.Db

  var err = db.Find(&dirs)
  if err != nil{
    fmt.Println(err.Error())
    panic("Couldn't find all the dirs")
  }

  for _, dir := range dirs{
    GetTodosInDirectory(&dir)
  }
}

func AddNewTodo(dir *models.Directory, newTodo *models.Todo){
  var db = connector.Db

  var has, err = db.Exist(dir)
  if err != nil{
    fmt.Println(err.Error())
    panic("Couldn't determeine if a directory exists")
  }
  if has == false{
    _ = AddNewDirectory(dir)
  }

  GetTodosInDirectory(dir)
  var old_todos = dir.Todos
  var new_todos = append(old_todos, *newTodo)
  dir.Todos = new_todos

  var affected, new_err = db.Cols("todos").Update(dir)
  if new_err != nil{
    fmt.Println(new_err.Error())
    panic("Couldn't update that record")
  }
  fmt.Println(affected)
}

func AddNewDirectory(dir *models.Directory) error{
  var db = connector.Db

  var _, err = db.Insert(dir)
  return err
}

func CloseConnection(){
  connector.CloseConn()
}
