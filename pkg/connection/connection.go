package connector;

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "models"
)

type DBConnection struct {
  db *DB
  err error
  connected bool
}

func (conn DBConnection) connect(db_path string) {
  if conn.connected == false{
    conn.db, conn.err = gorm.Open("sqlite3", db_path);
    conn.connected = true;
  }
  if conn.err != nil {
    panic("failed to connect to db");
  }
  defer conn.db.close();
}

func (conn DBConnection) close_conn(){
  conn.db.close();
}

func (conn DBConnection) migrate(){
  conn.db.AutoMigrate(&Directory{}, &TodoBase{}, &Todo{});
}
