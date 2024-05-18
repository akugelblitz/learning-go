package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
)

type saver interface {
  Save() error
}


func main(){
  title, content := getNoteData()
  firstnote, err := note.New(title, content)
  if err != nil {
    fmt.Println(err)
    return
  }
  firstnote.Log()
  err = saveData(firstnote)
  
}

func saveData(data saver) error{
  err := data.Save()
  if err != nil{
    fmt.Println("Saving failed: ", err)
    return err
  } else{
    fmt.Println("Successfully saved the file")
  }
 
  return nil
}

func getNoteData() (string, string){
  title := getUserInput("Enter title of the note: ")
  content := getUserInput("Enter content of the note: ")

  return title, content
}

func getUserInput(prompt string) (string){
  fmt.Println(prompt)

  var input string
  reader := bufio.NewReader(os.Stdin)
  input, err := reader.ReadString('\n')

  if err != nil{
    fmt.Println(err)
    return ""
  }
  input = strings.TrimSuffix(input, "\n")
  input = strings.TrimSuffix(input, "\r")
  return input
}
