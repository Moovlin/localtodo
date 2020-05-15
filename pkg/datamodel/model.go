package datamodel

import (
  "fmt"
  "github.com/moovlin/localtodo/pkg/connection"
  "github.com/moovlin/localtodo/pkg/models"
)

func OpenConnection(path string){
  connector.Connect(path)
}

func GetTodosInDirectory(dir models.Directory){
  var db = connector.Db

  var found, err = db.Get(&dir)
  if err != nil || found == false{
    fmt.Println(err.Error())
    panic("Couldn't get the table")
  }

  for _, todo := range dir.Todos{
    fmt.Println(todo.Text)
  }
}

func GetAllTodos(){
  var db = connector.Db

  var dirs []models.Directory
  var err = db.Find(&dirs)
  if err != nil{
    fmt.Println(err.Error())
    panic("Couldn't find all the dirs")
  }

  for _, dir := range dirs{
    GetTodosInDirectory(dir)
  }
}

func AddNewTodo(dir models.Directory, newTodo models.Todo){
  
}

func AddNewTodo(dir string, newTodo models.Todo){

}

func AddNewDirectory(){

}

func CloseConnection(){
  connector.CloseConn()
}
