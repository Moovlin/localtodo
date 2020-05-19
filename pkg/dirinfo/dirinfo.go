package dirsinfo


import(
  "fmt"
  "strings"
  "errors"
  "strconv"
  "io/ioutil"
  "encoding/base64"
  "encoding/binary"
  "path/filepath"
  "github.com/moovlin/localtodo/pkg/models"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+/"
var encoder = base64.NewEncoding(letterBytes)

func LoadID(dir models.Directory) error{
  var rootDir = dir.Path
  if rootDir == ""{
    rootDir = "."
  }
  matches, err := filepath.Glob(rootDir + "/.*.todo")
  if err != nil{
    fmt.Println(err.Error())
  }

  var thisError error
  if len(matches) > 1{
    thisError = errors.New("Too many todos in this directory. We only support one list at the moment. Using the first one")
  }

  if len(matches) == 0{
    thisError = errors.New("No matches")
  }

  var id = strings.TrimSuffix(matches[0], ".todo")
  id = strings.TrimPrefix(id, ".")
  fmt.Println(id)
  dir.Id, _  = strconv.ParseInt(id, 64, 64)
  return thisError
}

func SaveId(dir models.Directory) error{
  var rootDir = dir.Path
  if rootDir == "" {
    rootDir = "."
  }

  var intByteBuff = make([]byte, 8)
  _ = binary.PutVarint(intByteBuff, dir.Id)
  var saveId = encoder.EncodeToString(intByteBuff)
  var err = ioutil.WriteFile(rootDir + "." + saveId + ".todo", make([]byte, 0), 0644)
  return err
}
