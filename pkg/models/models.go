package models

import (
  "time"
)


type Todo struct {
  Id int64

  // ToDo information
  TodoLevel uint
  Text string `xorm:"notnull varchar(512)"`
  Completed bool
  Created time.Time `xorm:"created"`
  Updated time.Time `xorm:"updated"`

}

type Directory struct {
  Path string `xorm:"pk varchar(1024)"`
  Todos []Todo
}


