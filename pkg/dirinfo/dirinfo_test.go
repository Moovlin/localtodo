package dirsinfo

import (
  "testing"
  "fmt"
  "github.com/moovlin/localtodo/pkg/models"
  "github.com/moovlin/localtodo/pkg/datamodel"
  "github.com/moovlin/localtodo/pkg/connector"
)

func TestLoadID(t *testing.T){
  var dir models.Directory
  var test_dir = "/home/sinn3r"
  dir.Path = test_dir
  dir.Id = -1

  var _ = model.AddNewDirectory(&dir)
  if dir.Id == -1{
    t.Errorf("Id didn't update.")
  }

  if dir.Path != test_dir{
    t.Errorf("Testing dir didn't match up")
  }

  model.TrueDeleteDirectory(&dir)
}

// func TestGetAllTodos(t *testing.T){
//
// }
