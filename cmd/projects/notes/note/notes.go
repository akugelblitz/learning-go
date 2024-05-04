package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct{
  Title string `json:"title"`
  Content string `json:"content"`
  CreatedAt time.Time `json:"created_at"`
}

func New(title string, content string) (*Note, error){

  if title != "" || content != ""{
    note := &Note{
      Title : title,
      Content : content,
      CreatedAt : time.Now(),
    }

    return note, nil
  }else{
    return nil, errors.New("Title and content cannot be empty")
  }
}

func (note *Note) Log(){
  fmt.Printf("The note created at %v with the title %v has the following content: \n%v\n", note.CreatedAt, note.Title, note.Content)
}

func (note *Note) Save() error{
  fileName := strings.ReplaceAll(note.Title, " ", "_")
  fileName = strings.ToLower(fileName) + ".json"
  
  json, err := json.Marshal(note)

  if err != nil{
    return err
  }
  return os.WriteFile(fileName, json, 0644)

}

