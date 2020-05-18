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

  fmt.Println("Done with connection test\nStarting datamodel test")

  datamodel.OpenConnection(test_db)
  datamodel.GetTodosInDirectory(&dir)
  var dirs []models.Directory
  fmt.Println("Getting All Todos")
  datamodel.GetAllTodos(dirs)


  var newPath = "/home/sinn3r"
  var newDir = new(models.Directory)
  newDir.Path = newPath
  newDir.Todos = make([]models.Todo, 1)
  models.BaseTodo(newDir.Todos[0])


  var newTodo = new(models.Todo)
  newTodo.Text = "testing shit"
  datamodel.AddNewTodo(newDir, newTodo)

  var testingDir = new(models.Directory)
  datamodel.GetTodosInDirectory(testingDir)
  for _, todo := range testingDir.Todos {
    fmt.Println(todo.Text)
  }

  datamodel.CloseConnection()
}
