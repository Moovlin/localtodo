package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)


type Todo struct {
  gorm.Model

  // PKey
  TodoID []uint `gorm:"primary_key;AUTO_INCREMENT"`

  // ToDo information
  TodoLevel uint
  text string `gorm:"not null;size:512"`
  completed bool

}

type TodoBase struct {
  gorm.Model
  Todos []Todo  `gorm:"foreignkey:TodoID"`
  DirID uint  `gorm:"primary_key"`
}

type Directory struct {
  gorm.Model
  Path string `gorm:"primary_key"`
  Todos TodoBase `gorm:"foreignkey:DirID"`
}


