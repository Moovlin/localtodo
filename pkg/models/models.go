package models

import (
  "time"
)


type Todo struct {
  Id int64
  ParentId int64

  // ToDo information
  TodoLevel uint
  Text string `xorm:"notnull varchar(512)"`
  Completed bool
  Deleted bool
  Created time.Time `xorm:"created"`
  Updated time.Time `xorm:"updated"`

}

type Directory struct {
  Path string `xorm:"unique"`
  Id int64
  Todos []Todo
}

func BaseTodo(emptyTodo Todo){
  emptyTodo.Text = "BASE"
  emptyTodo.ParentId = 0
}
