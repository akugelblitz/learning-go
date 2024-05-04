package main

import "fmt"

// type user struct {
//   firstName string
//   lastName string
//   dob string
// }
// func (appUser user) outputUserData(){
//   fmt.Println(appUser.firstName, appUser.lastName, appUser.dob)
// }
//
// func main() {
//   var appUser user = user{}
//   appUser.firstName = "hi"
//   appUser.firstName = getUserData("Enter first name")
//   appUser.lastName = getUserData("Enter last name")
//   appUser.dob = getUserData("Enter dob")
//
//   appUser.outputUserData()
// }
//
//
// func getUserData(prompt string) string {
//   fmt.Println(prompt)
//   var inp string
//   fmt.Scan(&inp)
//   return inp
// }
//
//

type str string

func (text str) log() {
	fmt.Println(text)
}
func main() {
	var name str = "hi"
	name.log()
}
